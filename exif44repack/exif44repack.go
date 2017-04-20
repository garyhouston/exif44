package main

import (
	"bytes"
	"errors"
	"fmt"
	exif "github.com/garyhouston/exif44"
	jseg "github.com/garyhouston/jpegsegs"
	tiff "github.com/garyhouston/tiff66"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func processTIFF(file io.Reader) ([]byte, error) {
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	validTIFF, order, ifdPos := tiff.GetHeader(buf)
	if !validTIFF {
		return nil, errors.New("processTIFF: invalid TIFF header")
	}
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

func processJPEG(file io.Reader) ([]byte, error) {
	scanner, err := jseg.NewScanner(file)
	if err != nil {
		return nil, err
	}
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

// Decode a TIFF file, or the Exif segment in a JPEG file, then re-encode
// it and write to a new file.
func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s file outfile\n", os.Args[0])
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
	var out []byte
	if fileType == TIFFFile {
		out, err = processTIFF(file)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		out, err = processJPEG(file)
		if err != nil {
			log.Fatal(err)
		}
	}
	ioutil.WriteFile(os.Args[2], out, 0644)
}
