package main

import (
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

func processTIFF(infile io.Reader, outfile io.Writer) error {
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
	root.Fix(order)
	fileSize := tiff.HeaderSize + root.TreeSize()
	out := make([]byte, fileSize)
	tiff.PutHeader(out, order, tiff.HeaderSize)
	_, err = root.PutIFDTree(out, tiff.HeaderSize, order)
	if err != nil {
		return err
	}
	_, err = outfile.Write(out)
	return err
}

func processJPEG(infile io.Reader, outfile io.Writer) error {
	scanner, err := jseg.NewScanner(infile)
	if err != nil {
		return err
	}
	dumper, err := jseg.NewDumper(outfile)
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
			if err := dumper.Dump(marker, nil); err != nil {
				return err
			}
			err := dumper.Copy(scanner)
			return err
		}
		if marker == jseg.APP0+1 {
			isExif, next := exif.GetHeader(buf)
			if isExif {
				tree, err := exif.GetExifTree(buf[next:])
				if err != nil {
					return err
				}
				tree.Tree.Fix(tree.Order)
				app1 := make([]byte, exif.HeaderSize+tree.TreeSize())
				next := exif.PutHeader(app1)
				_, err = tree.Put(app1[next:])
				if err != nil {
					return err
				}
				buf = app1
			}

		}
		if err := dumper.Dump(marker, buf); err != nil {
			return err
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
	infile, err := os.Open(os.Args[1])
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
	outfile, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	if fileType == TIFFFile {
		err = processTIFF(infile, outfile)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = processJPEG(infile, outfile)
		if err != nil {
			log.Fatal(err)
		}
	}
}
