package main

import (
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	exif "github.com/garyhouston/exif44"
	jseg "github.com/garyhouston/jpegsegs"
	tiff "github.com/garyhouston/tiff66"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func printTree(node *tiff.IFDNode, order binary.ByteOrder, length uint32) {
	fmt.Println()
	fields := node.Fields
	fmt.Printf("%s IFD with %d ", node.Space.Name(), len(fields))
	if len(fields) > 1 {
		fmt.Println("entries:")
	} else {
		fmt.Println("entry:")
	}
	names := exif.TagNameMap(node.Space)
	for i := 0; i < len(fields); i++ {
		fields[i].Print(order, names, length)
	}
	for i := 0; i < len(node.SubIFDs); i++ {
		printTree(node.SubIFDs[i].Node, order, length)
	}
	if node.Next != nil {
		printTree(node.Next, order, length)
	}
}

func scanTIFF(buf []byte, maxLen uint32) error {
	validTIFF, order, ifdPos := tiff.GetHeader(buf)
	if !validTIFF {
		return errors.New("scanTIFF: invalid TIFF header")
	}
	root, err := tiff.GetIFDTree(buf, order, ifdPos, tiff.TIFFSpace)
	if err != nil {
		return err
	}
	printTree(root, order, maxLen)
	return nil
}

func processTIFF(file io.Reader, maxLen uint32) error {
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	return scanTIFF(buf, maxLen)
}

// Process a single image in a JPEG file. A file using the
// Multi-Picture Format extension will contain multiple images.
func processImage(reader io.ReadSeeker, maxLen uint32, mpfProcessor jseg.MPFProcessor) error {
	scanner, err := jseg.NewScanner(reader)
	if err != nil {
		return err
	}
	for {
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
				if err := scanTIFF(buf[next:], maxLen); err != nil {
					return err
				}
			}
		}
		if marker == jseg.APP0+2 {
			_, _, err := mpfProcessor.ProcessAPP2(nil, reader, buf)
			if err != nil {
				return err
			}
		}
	}
}

// State for the MPF image iterator.
type scanData struct {
	maxLen uint32
}

// Function to be applied to each MPF image.
func (scan *scanData) MPFApply(reader io.ReadSeeker, index uint32, length uint32) error {
	if index > 0 {
		fmt.Println()
		fmt.Println("== Processing Image ", index+1, "==")
		return processImage(reader, scan.maxLen, &jseg.MPFCheck{})
	}
	return nil
}

// Process a JPEG file.
func processJPEG(file io.ReadSeeker, maxLen uint32) error {
	var index jseg.MPFGetIndex
	if err := processImage(file, maxLen, &index); err != nil {
		return err
	}
	if index.Index != nil {
		scandata := &scanData{maxLen}
		err := index.Index.ImageIterate(file, scandata)
		if err != nil {
			log.Fatal(err)
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

// Read and diplay all the IFDs of a TIFF or Exif segment of a JPEG
// file, including any private IFDs that can be detected.
func main() {
	var maxLen uint
	flag.UintVar(&maxLen, "m", 20, "maximum values to print or 0 for no limit")
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Printf("Usage: %s [-m max values] file\n", os.Args[0])
		return
	}
	file, err := os.Open(flag.Arg(0))
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
		if err := processTIFF(file, uint32(maxLen)); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := processJPEG(file, uint32(maxLen)); err != nil {
			log.Fatal(err)
		}
	}
}
