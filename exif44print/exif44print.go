package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	exif "github.com/garyhouston/exif44"
	jseg "github.com/garyhouston/jpegsegs"
	tiff "github.com/garyhouston/tiff66"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func printNode(node *tiff.IFDNode, order binary.ByteOrder, length uint32) {
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
		printNode(node.SubIFDs[i].Node, order, length)
	}
	if node.Next != nil {
		printNode(node.Next, order, length)
	}
}

func processJPEG(scanner *jseg.Scanner, maxLen uint32) error {
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
				printNode(exif.Tree, exif.Order, maxLen)
			}
		}
	}
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
	buf, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	validTIFF, order, ifdPos := tiff.GetHeader(buf)
	if validTIFF {
		root, err := tiff.GetIFDTree(buf, order, ifdPos, tiff.TIFFSpace)
		if err != nil {
			log.Fatal(err)
		}
		printNode(root, order, uint32(maxLen))
	} else {
		reader := strings.NewReader(string(buf))
		scanner, err := jseg.NewScanner(reader)
		if err != nil {
			log.Fatal(flag.Arg(0), " not a valid TIFF or JPEG file")
		}
		err = processJPEG(scanner, uint32(maxLen))
		if err != nil {
			log.Fatal(err)
		}
	}
}
