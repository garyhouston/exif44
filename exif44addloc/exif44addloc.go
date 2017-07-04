package main

// Add location coordinates to a JPEG or TIFF file.

import (
	"errors"
	"encoding/binary"
	exif "github.com/garyhouston/exif44"
	"fmt"
	"math"
	jseg "github.com/garyhouston/jpegsegs"
	tiff "github.com/garyhouston/tiff66"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// Insert location coordinates into a GPS IFD.
func insertGPS(lat, long float64, gpsNode *tiff.IFDNode) {
	var latRef, longRef string
	if lat < 0 {
		latRef = "S"
		lat = -lat
	} else {
		latRef = "N"
	}
	if long < 0 {
		longRef = "W"
		long = -long
	} else {
		longRef = "E"
	}
	// Arbitarily setting 6 decimal places in seconds fields.
	secMult := uint32(1000000)
	latDeg := uint32(lat)
	latMin := uint32(math.Mod(lat * 60, 60))
	latSec := uint32(math.Mod(lat * 3600, 60)*float64(secMult))
	longDeg := uint32(long)
	longMin := uint32(math.Mod(long * 60, 60))
	longSec := uint32(math.Mod(long * 3600, 60)*float64(secMult))
	latData := make([]byte, tiff.RATIONAL.Size() * 3)
	longData := make([]byte, tiff.RATIONAL.Size() * 3)
	fields := make([]tiff.Field, 4)
	fields[0] = tiff.Field{exif.GPSLatitudeRef, tiff.ASCII, 2, []byte(latRef)}
	fields[1] = tiff.Field{exif.GPSLongitudeRef, tiff.ASCII, 2, []byte(longRef)}
	fields[2] = tiff.Field{exif.GPSLatitude, tiff.RATIONAL, 3, latData}
	fields[3] = tiff.Field{exif.GPSLongitude, tiff.RATIONAL, 3, longData}
	fields[2].PutRational(latDeg, 1, 0, gpsNode.Order)
	fields[2].PutRational(latMin, 1, 1, gpsNode.Order)
	fields[2].PutRational(latSec, secMult, 2, gpsNode.Order)
	fields[3].PutRational(longDeg, 1, 0, gpsNode.Order)
	fields[3].PutRational(longMin, 1, 1, gpsNode.Order)
	fields[3].PutRational(longSec, secMult, 2, gpsNode.Order)
	gpsNode.IFD_T.DeleteFields([]tiff.Tag{exif.GPSLatitudeRef, exif.GPSLatitude, exif.GPSLongitudeRef, exif.GPSLongitude})
	gpsNode.IFD_T.AddFields(fields)
}

// Add a GPS sub-IFD to an IFD node and return a pointer to it.
func addGPS(node *tiff.IFDNode) *tiff.IFDNode {
	gpsNode := new(tiff.IFDNode)
	gpsNode.Order = node.Order
	gpsNode.Space = tiff.GPSSpace
	gpsVersionData := make([]byte, 4)
	gpsVersion := tiff.Field{exif.GPSVersionID, tiff.BYTE, 4, gpsVersionData}
	gpsVersion.PutByte(2, 0)
	gpsVersion.PutByte(3, 1)
	gpsNode.IFD_T.AddFields([]tiff.Field{gpsVersion})
	gpsIFDData := make([]byte, 4)
	node.IFD_T.AddFields([]tiff.Field{{tiff.GPSIFD, tiff.LONG, 1, gpsIFDData}})
	subIFD := tiff.SubIFD{tiff.GPSIFD, gpsNode}
	node.SubIFDs = append(node.SubIFDs, subIFD)
	return gpsNode
}

// Add location coordinates into a TIFF tree, creating a GPS IFD if needed.
func putGPS(lat, long float64, root *tiff.IFDNode) {
	var gpsNode *tiff.IFDNode
	for _, i := range root.SubIFDs {
		if i.Tag == tiff.GPSIFD {
			gpsNode = i.Node
		}
	}
	if gpsNode == nil {
		gpsNode = addGPS(root)
	}
	insertGPS(lat, long, gpsNode)
}

// Process a TIFF file.
func processTIFF(lat, long float64, outfile io.Writer, infile io.Reader) error {
	buf, err := ioutil.ReadAll(infile)
	if err != nil {
		return err
	}
	validTIFF, order, ifdPos := tiff.GetHeader(buf)
	if !validTIFF {
		return errors.New("processTIFF: invalid TIFF header")
	}
	root, err := tiff.GetIFDTree(buf, order, ifdPos, tiff.TIFFSpace)
	if err != nil {
		return err
	}
	root.Fix()
	putGPS(lat, long, root)
	fileSize := tiff.HeaderSize + root.TreeSize()
	out := make([]byte, fileSize)
	tiff.PutHeader(out, order, tiff.HeaderSize)
	_, err = root.PutIFDTree(out, tiff.HeaderSize)
	if err != nil {
		return err
	}
	_, err = outfile.Write(out)
	return err
}

// Create a new Exif node for a JPEG image, from scratch.
func createExif() *exif.Exif {
	node := new(tiff.IFDNode)
	node.Order = binary.LittleEndian // arbitrary
	node.Space = tiff.ExifSpace
	// Add an Exif sub-IFD too, since the Exif spec may require
	// the version to be specified.
	exifNode := new(tiff.IFDNode)
	exifNode.Order = node.Order
	exifNode.Space = tiff.ExifSpace
	exifVersionData := make([]byte, 4)
	copy(exifVersionData, []byte("0230"))
	exifVersion := tiff.Field{exif.ExifVersion, tiff.UNDEFINED, 4, exifVersionData}
	exifNode.IFD_T.AddFields([]tiff.Field{exifVersion})
	exifIFDData := make([]byte, 4)
	node.IFD_T.AddFields([]tiff.Field{{tiff.ExifIFD, tiff.LONG, 1, exifIFDData}})
	subIFD := tiff.SubIFD{tiff.ExifIFD, exifNode}
	node.SubIFDs = append(node.SubIFDs, subIFD)
	return &exif.Exif{TIFF: node}
}

// Process a single image in a JPEG file. A file using Multi-Picture
// Format will contain multiple images.
func processImage(writer io.WriteSeeker, reader io.ReadSeeker, index uint32, lat, long float64, mpfProcessor jseg.MPFProcessor) error {
	scanner, err := jseg.NewScanner(reader)
	if err != nil {
		return err
	}
	dumper, err := jseg.NewDumper(writer)
	if err != nil {
		return err
	}
	segs, err := jseg.ReadSegments(scanner)
	if err != nil {
		return err
	}
	var exifIdx int
	var exifNode *exif.Exif
	for idx, seg := range segs {
		// If multiple images are present, add GPS coordinates
		// only to the first.
		if index == 0 && seg.Marker == jseg.APP0+1 {
			isExif, next := exif.GetHeader(seg.Data)
			if isExif {
				exifIdx = idx
				exifNode, err = exif.GetExifTree(seg.Data[next:])
				if err != nil {
					return err
				}
				exifNode.TIFF.Fix()
			}
		}
		if seg.Marker == jseg.APP0+2 {
			_, buf, err := mpfProcessor.ProcessAPP2(writer, reader, seg.Data)
			if err != nil {
				return err
			}
			seg.Data = buf
		}
	}
	if index == 0 && exifNode == nil {
		// No Exif block in the file, so we'll have to create one.
		exifNode = createExif()
		// Make it the first segment.
		newSeg := jseg.Segment{jseg.APP0+1, nil}
		segs = append([]jseg.Segment{newSeg}, segs...)
		exifIdx = 0
	}
	if exifNode != nil {
		putGPS(lat, long, exifNode.TIFF)
		app1 := make([]byte, exif.HeaderSize+exifNode.TreeSize())
		next := exif.PutHeader(app1)
		_, err = exifNode.Put(app1[next:])
		if err != nil {
			return err
		}
		segs[exifIdx].Data = app1
	}
	if err = jseg.WriteSegments(dumper, segs); err != nil {
		return err
	}
	// Process the rest of the file, image data expected.
	for {
		marker, buf, err := scanner.Scan()
		if err != nil {
			return err
		}
		if err := dumper.Dump(marker, buf); err != nil {
			return err
		}
		if marker == jseg.EOI {
			return nil
		}
	}
}

// State for MPF image iterator.
type iterData struct {
	writer     io.WriteSeeker
	newOffsets []uint32
	lat float64
	long float64
}

// Function to be applied to each MPF image.
func (iter *iterData) MPFApply(reader io.ReadSeeker, index uint32, length uint32) error {
	if index > 0 {
		pos, err := iter.writer.Seek(0, io.SeekCurrent)
		if err != nil {
			return err
		}
		iter.newOffsets[index] = uint32(pos)
		return processImage(iter.writer, reader, index, iter.lat, iter.long, &jseg.MPFCheck{})
	}
	return nil
}

// Process additional images found in the MPF index.
func processMPFImages(writer io.WriteSeeker, reader io.ReadSeeker, lat, long float64, index *jseg.MPFIndex) ([]uint32, error) {
	var iter iterData
	iter.writer = writer
	iter.newOffsets = make([]uint32, len(index.ImageOffsets))
	iter.lat = lat
	iter.long = long
	index.ImageIterate(reader, &iter)
	return iter.newOffsets, nil
}

// Process a JPEG file.
func processJPEG(lat, long float64, writer io.WriteSeeker, reader io.ReadSeeker) error {
	var mpfIndex jseg.MPFIndexRewriter
	if err := processImage(writer, reader, 0, lat, long, &mpfIndex); err != nil {
		return err
	}
	if mpfIndex.Tree != nil {
		newOffsets, err := processMPFImages(writer, reader, lat, long, mpfIndex.Index)
		if err != nil {
			return err
		}
		end, err := writer.Seek(0, io.SeekCurrent)
		if err != nil {
			return err
		}
		if err = jseg.RewriteMPF(writer, mpfIndex.Tree, mpfIndex.APP2WritePos, newOffsets, uint32(end)); err != nil {
			return err
		}
	}
	return nil
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

func usage() {
	fmt.Printf("Usage: %s latitude longitude file outfile\nLatitude and longitude are expressed as signed decimals\n", os.Args[0])
}

// Add location coordinates to a TIFF or JPG file.
func main() {
	if len(os.Args) != 5 {
		usage()
		return
	}
	latStr := os.Args[1]
	longStr := os.Args[2]
	in := os.Args[3]
	out := os.Args[4]
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		usage()
		return
	}
	if lat < -90 || lat > 90 {
		log.Fatal("Lattitude is out of range [-90, 90]")
	}
	long, err := strconv.ParseFloat(longStr, 64)
	if err != nil {
		usage()
		return
	}
	if long < -180 || long > 180 {
		log.Fatal("Longitude is out of range [-180, 180]")
	}
	infile, err := os.Open(in)
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()
	fileType, err := fileType(infile)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := infile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}
	outfile, err := os.Create(out)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	if fileType == TIFFFile {
		err = processTIFF(lat, long, outfile, infile)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = processJPEG(lat, long, outfile, infile)
		if err != nil {
			log.Fatal(err)
		}
	}
}
