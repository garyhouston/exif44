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
	fields := node.IFD.Fields
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

func processTIFF(file io.Reader, maxLen uint32) error {
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	validTIFF, order, ifdPos := tiff.GetHeader(buf)
	if !validTIFF {
		return errors.New("processTIFF: not a TIFF file")
	}
	root, err := tiff.GetIFDTree(buf, order, ifdPos, tiff.TIFFSpace)
	if err != nil {
		return err
	}
	printTree(root, order, maxLen)
	return nil
}

func processJPEG(file io.Reader, maxLen uint32) error {
	scanner, err := jseg.NewScanner(file)
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
				exif, err := exif.GetExifTree(buf[next:])
				if err != nil {
					return err
				}
				printTree(exif.Tree, exif.Order, maxLen)
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
