package main

import (
	"flag"
	"fmt"
	exif "github.com/garyhouston/exif44"
	tiff "github.com/garyhouston/tiff66"
	"log"
	"os"
)

// Recursively print an IFD node, its subIFDs, and next IFD.
func printTree(format exif.FileFormat, node *tiff.IFDNode, maxLen uint32) {
	fmt.Println()
	fields := node.Fields
	space := node.GetSpace()
	fmt.Printf("%s IFD with %d ", space.Name(), len(fields))
	if len(fields) != 1 {
		fmt.Println("entries:")
	} else {
		fmt.Println("entry:")
	}
	order := node.Order
	names := exif.TagNameMap(space)
	for i := 0; i < len(fields); i++ {
		fields[i].Print(order, names, maxLen)
	}
	for i := 0; i < len(node.SubIFDs); i++ {
		printTree(format, node.SubIFDs[i].Node, maxLen)
	}
	if format == exif.FileJPEG && node.Next != nil {
		printTree(format, node.Next, maxLen)
	}
}

// Exif handler.
type readExif struct {
	maxLen uint32
}

func (readExif readExif) ReadExif(format exif.FileFormat, imageIdx uint32, exif exif.Exif) error {
	if imageIdx > 0 {
		fmt.Println()
		fmt.Println("== Processing Image ", imageIdx+1, "==")
	}
	printTree(format, exif.TIFF, readExif.maxLen)
	return nil
}

// Read and print all the IFDs of a TIFF file, or Exif segment of a
// JPEG file, including any private IFDs that can be detected.
func main() {
	var maxLen uint
	flag.UintVar(&maxLen, "m", 20, "maximum values to print or 0 for no limit")
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Printf("Usage: %s [-m max values] file\n", os.Args[0])
		return
	}
	var control exif.ReadControl
	control.ReadExif = readExif{maxLen: uint32(maxLen)}
	if err := exif.ReadFile(flag.Arg(0), control); err != nil {
		log.Fatal(err)
	}
}
