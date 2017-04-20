package main

import (
	"bytes"
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

func processTIFF(buf []byte, order binary.ByteOrder, ifdPos uint32) ([]byte, error) {
	root, err := tiff.GetIFDTree(buf, order, ifdPos, tiff.TIFFSpace)
	if err != nil {
		return nil, err
	}
	root.Fix(order)
	fileSize := tiff.HeaderSize + root.TreeSize()
	out := make([]byte, fileSize)
	tiff.PutHeader(out, order, tiff.HeaderSize)
	next, err := root.PutIFDTree(out, tiff.HeaderSize, order)
	if err != nil {
		return nil, err
	}
	out = out[:next]
	return out, nil
}

func processJPEG(scanner *jseg.Scanner) ([]byte, error) {
	out := new(bytes.Buffer)
	dumper, err := jseg.NewDumper(out)
	if err != nil {
		return out.Bytes(), err
	}
	for {
		marker, buf, err := scanner.Scan()
		if err != nil {
			return out.Bytes(), err
		}
		if marker == jseg.SOS {
			// Start of scan data, no more metadata expected.
			if err := dumper.Dump(marker, nil); err != nil {
				return out.Bytes(), err
			}
			err := dumper.Copy(scanner)
			return out.Bytes(), err
		}
		if marker == jseg.APP0+1 {
			isExif, next := exif.GetHeader(buf)
			if isExif {
				tree, err := exif.GetExifTree(buf[next:])
				if err != nil {
					return out.Bytes(), err
				}
				tree.Tree.Fix(tree.Order)
				app1 := make([]byte, exif.HeaderSize+tree.TreeSize())
				next := exif.PutHeader(app1)
				_, err = tree.Put(app1[next:])
				if err != nil {
					return out.Bytes(), err
				}
				buf = app1
			}

		}
		if err := dumper.Dump(marker, buf); err != nil {
			return out.Bytes(), err
		}
	}
}

// Decode a TIFF file, or the Exif segment in a JPEG file, then re-encode
// it and write to a new file.
func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s file outfile\n", os.Args[0])
		return
	}
	buf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	validTIFF, order, ifdPos := tiff.GetHeader(buf)
	var out []byte
	if validTIFF {
		var err error
		out, err = processTIFF(buf, order, ifdPos)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		reader := strings.NewReader(string(buf))
		scanner, err := jseg.NewScanner(reader)
		if err != nil {
			log.Fatal(flag.Arg(0), " not a valid TIFF or JPEG file")
		}
		out, err = processJPEG(scanner)
		if err != nil {
			log.Fatal(err)
		}
	}
	ioutil.WriteFile(os.Args[2], out, 0644)
}
