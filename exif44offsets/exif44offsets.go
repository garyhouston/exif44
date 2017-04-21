package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	exif "github.com/garyhouston/exif44"
	jseg "github.com/garyhouston/jpegsegs"
	tiff "github.com/garyhouston/tiff66"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

// This program will print the byte layout of IFDs (image file
// directories) and related data in a TIFF file, or in the Exif block
// of a JPEG file. If the TIFF data doesn't start at the beginning of
// the file, it will print a header with the offset, and remaining
// positions will be relative to the start of the TIFF header. Each
// line of output will have a start and ending position and a
// description of where the data originated. If the description is
// blank, nothing is known and it may a data region for a field not
// supported by the tiff66 library. Unused single bytes are likely to
// be fillers to align the next IFD. Maker notes are included in IFD
// external data.

// How bytes in a TIFF file are used.
type usage int

const (
	usageHeader    = 1 // TIFF header.
	usageIFD       = 2 // Part of an IFD.
	usageData      = 3 // IFD external data.
	usageImageData = 4 // IFD image data, or other data referred to by pointer.
)

type position struct {
	start  uint32
	size   uint32
	IFDpos uint32
	space  tiff.TagSpace
	usage  usage
}

// Combine continguous external data for an IFD into a single block.
func combine(positions []position) []position {
	count := len(positions)
	i := 0
	for {
		if i+1 == count {
			break
		}
		if (positions[i].usage == usageData || positions[i].usage == usageImageData) &&
			positions[i].usage == positions[i+1].usage &&
			positions[i].IFDpos == positions[i+1].IFDpos &&
			positions[i].start+positions[i].size == positions[i+1].start {
			positions[i].size += positions[i+1].size
			positions = append(positions[:i+1], positions[i+2:]...)
			count--
		} else {
			i++
		}
	}
	return positions
}

func print(offset int64, positions []position, bufLen uint32) {
	sort.Slice(positions, func(i, j int) bool { return positions[i].start < positions[j].start })
	positions = combine(positions)
	next := uint32(0)
	if offset > 0 {
		fmt.Printf("Start of TIFF at position %d in file\n", offset)
	}
	for _, pos := range positions {
		start := pos.start
		end := pos.start + pos.size - 1
		if start != next {
			fmt.Printf("%9d %9d\n", next, start-1)
		}
		next = end + 1
		switch pos.usage {
		case usageHeader:
			fmt.Printf("%9d %9d TIFF header", start, end)
		case usageIFD:
			fmt.Printf("%9d %9d %s IFD", start, end, pos.space.Name())
		case usageData:
			fmt.Printf("%9d %9d", start, end)
			fmt.Printf(" external data for IFD at %d", pos.IFDpos)
		case usageImageData:
			fmt.Printf("%9d %9d", start, end)
			fmt.Printf(" image data for IFD at %d", pos.IFDpos)
		default:
			fmt.Printf("invalid entry in positions data")
		}
		fmt.Println()
	}
	if next < bufLen {
		fmt.Printf("%9d %9d\n", next, bufLen-1)
	}
}

// Scan a node in an IFD tree recursively.
func scanTree(buf []byte, order binary.ByteOrder, ifdPos uint32, space tiff.TagSpace, positions *[]position) error {
	fieldproc := func(idx uint16, field *tiff.Field, pos uint32) {
		size := field.Size()
		if size > 4 {
			*positions = append(*positions, position{pos, size, ifdPos, space, usageData})
		}
	}
	ifd, next, err := tiff.GetIFD(buf, order, ifdPos, tiff.TIFFImageData, fieldproc)
	if err != nil {
		return err
	}
	*positions = append(*positions, position{ifdPos, ifd.Size(), ifdPos, space, usageIFD})
	for i := range ifd.ImageData {
		offsetfield := ifd.FindFields([]tiff.Tag{ifd.ImageData[i].OffsetTag})[0]
		for j := range ifd.ImageData[i].Segments {
			offset := uint32(offsetfield.AnyInteger(uint32(j), order))
			size := uint32(len(ifd.ImageData[i].Segments[j]))
			*positions = append(*positions, position{offset, size, ifdPos, space, usageImageData})
		}
	}
	for _, field := range ifd.Fields {
		if field.IsIFD(space) {
			for j := uint32(0); j < field.Count; j++ {
				subspace := tiff.SubSpace(space, field.Tag)
				err := scanTree(buf, order, field.Long(j, order), subspace, positions)
				if err != nil {
					return err
				}
			}
		}
	}
	if next != 0 {
		var nspace tiff.TagSpace
		if space == tiff.ExifSpace {
			// The next IFD after an Exif IFD is a thumbnail
			// encoded as TIFF.
			nspace = tiff.TIFFSpace
		} else {
			// Assume the next IFD is the same type.
			nspace = space
		}
		err = scanTree(buf, order, next, nspace, positions)
		if err != nil {
			return err
		}
	}
	return nil
}

func scanTIFF(buf []byte, offset int64) error {
	validTIFF, order, ifdPos := tiff.GetHeader(buf)
	if !validTIFF {
		return errors.New("scanTIFF: invalid TIFF header")
	}
	positions := make([]position, 1, 50)
	positions[0] = position{0, tiff.HeaderSize, 0, tiff.TIFFSpace, usageHeader}
	scanTree(buf, order, ifdPos, tiff.TIFFSpace, &positions)
	print(offset, positions, uint32(len(buf)))
	return nil
}

func processTIFF(file io.Reader) error {
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	return scanTIFF(buf, 0)
}

func processJPEG(file io.ReadSeeker) error {
	scanner, err := jseg.NewScanner(file)
	if err != nil {
		return err
	}
	for {
		offset, err := file.Seek(0, io.SeekCurrent)
		if err != nil {
			return err
		}
		marker, buf, err := scanner.Scan()
		if err != nil {
			return err
		}
		if marker == jseg.SOS {
			// Start of scan data, no more metadata expected.
			return nil
		}
		if marker == jseg.APP0+1 {
			isExif, next := exif.GetHeader(buf)
			if isExif {
				scanTIFF(buf[next:], offset+exif.HeaderSize)
			}
		}
	}
}

const (
	TIFFFile = 1
	JPEGFile = 2
)

// Determine if file is TIFF, JPEG or neither (error)
func fileType(file io.Reader) (int, error) {
	buf := make([]byte, tiff.HeaderSize)
	if _, err := io.ReadFull(file, buf); err != nil {
		return 0, err
	}
	if jseg.IsJPEGHeader(buf) {
		return JPEGFile, nil
	}
	if validTIFF, _, _ := tiff.GetHeader(buf); validTIFF {
		return TIFFFile, nil
	}
	return 0, errors.New("File doesn't have a TIFF or JPEG header")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s file\n", os.Args[0])
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileType, err := fileType(file)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := file.Seek(0, 0); err != nil {
		log.Fatal(err)
	}
	if fileType == TIFFFile {
		if err := processTIFF(file); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := processJPEG(file); err != nil {
			log.Fatal(err)
		}
	}
}
