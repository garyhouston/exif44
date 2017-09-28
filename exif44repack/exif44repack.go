package main

import (
	"fmt"
	exif "github.com/garyhouston/exif44"
	//	tiff "github.com/garyhouston/tiff66"
	"log"
	"os"
)

// Exif handler.
type readWriteExif struct {
}

func (readWriteExif readWriteExif) ReadWriteExif(imageIdx uint32, xif exif.Exif) (exif.Exif, error) {
	/*
	           // For the first image in the file, delete any "Software"
	           // field from the TIFF IFD0, and if there's an Exif IFD add a
	           // LensMake field.
	   	if imageIdx == 0 {
	   		xif.TIFF.DeleteFields([]tiff.Tag{tiff.Software})
	   		if xif.Exif != nil {
	   			xif.Exif.DeleteFields([]tiff.Tag{exif.LensMake})
	   			lens := "Dog Nose Lens"
	   			lensField := tiff.Field{
	   				Tag: exif.LensMake,
	   				Type: tiff.ASCII,
	   				Count: uint32(len(lens)),
	   				Data: make([]byte, len(lens) + 1)}
	   			lensField.PutASCII(lens)
	   			xif.Exif.AddFields([]tiff.Field{lensField})
	   		}
	   	}
	*/
	return xif, nil
}

// Decode a TIFF file, or the Exif segment(s) in a JPEG file, then re-encode
// it and write to a new file.
func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s file outfile\n", os.Args[0])
		return
	}
	var control exif.ReadWriteControl
	control.ReadWriteExif = readWriteExif{}
	if err := exif.ReadWriteFile(os.Args[1], os.Args[2], control); err != nil {
		log.Fatal(err)
	}
}
