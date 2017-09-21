package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	exif "github.com/garyhouston/exif44"
	tiff "github.com/garyhouston/tiff66"
	"log"
	"os"
)

// Recursively print an IFD node, its subIFDs, and next IFD.
func printTree(node *tiff.IFDNode, order binary.ByteOrder, maxLen uint32) {
	fmt.Println()
	fields := node.Fields
	space := node.GetSpace()
	fmt.Printf("%s IFD with %d ", space.Name(), len(fields))
	if len(fields) != 1 {
		fmt.Println("entries:")
	} else {
		fmt.Println("entry:")
	}
	names := exif.TagNameMap(space)
	for i := 0; i < len(fields); i++ {
		fields[i].Print(order, names, maxLen)
	}
	for i := 0; i < len(node.SubIFDs); i++ {
		printTree(node.SubIFDs[i].Node, order, maxLen)
	}
	if node.Next != nil {
		printTree(node.Next, order, maxLen)
	}
}

// Exif handler.
type readExif struct {
	maxLen uint32
}

func (readExif readExif) ReadExif(imageIdx uint32, exif exif.Exif) error {
	if imageIdx > 0 {
		fmt.Println()
		fmt.Println("== Processing Image ", imageIdx+1, "==")
	}
	printTree(exif.TIFF, exif.TIFF.Order, readExif.maxLen)
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
