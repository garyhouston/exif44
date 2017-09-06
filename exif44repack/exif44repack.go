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

// Process a TIFF file.
func processTIFF(outfile io.Writer, infile io.Reader) error {
	buf, err := ioutil.ReadAll(infile)
	if err != nil {
		return err
	}
	tree, err := exif.GetExifTree(buf)
	if err != nil {
		return err
	}
	tree.TIFF.Fix()
	if err = tree.CheckMakerNote(); err != nil {
		return err
	}
	if err = tree.MakerNoteComplexities(); err != nil {
		return err
	}
	fileSize := tiff.HeaderSize + tree.TreeSize()
	out := make([]byte, fileSize)
	tiff.PutHeader(out, tree.TIFF.Order, tiff.HeaderSize)
	_, err = tree.TIFF.PutIFDTree(out, tiff.HeaderSize)
	if err != nil {
		return err
	}
	_, err = outfile.Write(out)
	return err
}

// Process a single image in a JPEG file. A file using Multi-Picture
// Format will contain multiple images.
func processImage(writer io.WriteSeeker, reader io.ReadSeeker, mpfProcessor jseg.MPFProcessor) error {
	scanner, err := jseg.NewScanner(reader)
	if err != nil {
		return err
	}
	dumper, err := jseg.NewDumper(writer)
	if err != nil {
		return err
	}
	for {
		marker, buf, err := scanner.Scan()
		if err != nil {
			return err
		}
		if marker == jseg.APP0+1 {
			isExif, next := exif.GetHeader(buf)
			if isExif {
				tree, err := exif.GetExifTree(buf[next:])
				if err != nil {
					return err
				}
				tree.TIFF.Fix()
				if err = tree.CheckMakerNote(); err != nil {
					return err
				}
				if err = tree.MakerNoteComplexities(); err != nil {
					return err
				}
				app1 := make([]byte, exif.HeaderSize+tree.TreeSize())
				next := exif.PutHeader(app1)
				next, err = tree.Put(app1[next:])
				if err != nil {
					return err
				}
				buf = app1
			}

		}
		if marker == jseg.APP0+2 {
			_, buf, err = mpfProcessor.ProcessAPP2(writer, reader, buf)
			if err != nil {
				return err
			}
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
}

// Function to be applied to each MPF image.
func (iter *iterData) MPFApply(reader io.ReadSeeker, index uint32, length uint32) error {
	if index > 0 {
		pos, err := iter.writer.Seek(0, io.SeekCurrent)
		if err != nil {
			return err
		}
		iter.newOffsets[index] = uint32(pos)
		return processImage(iter.writer, reader, &jseg.MPFCheck{})
	}
	return nil
}

// Process additional images found in the MPF index.
func processMPFImages(writer io.WriteSeeker, reader io.ReadSeeker, index *jseg.MPFIndex) ([]uint32, error) {
	var iter iterData
	iter.writer = writer
	iter.newOffsets = make([]uint32, len(index.ImageOffsets))
	index.ImageIterate(reader, &iter)
	return iter.newOffsets, nil
}

// Process a JPEG file.
func processJPEG(writer io.WriteSeeker, reader io.ReadSeeker) error {
	var mpfIndex jseg.MPFIndexRewriter
	if err := processImage(writer, reader, &mpfIndex); err != nil {
		return err
	}
	if mpfIndex.Tree != nil {
		newOffsets, err := processMPFImages(writer, reader, mpfIndex.Index)
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
		err = processTIFF(outfile, infile)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = processJPEG(outfile, infile)
		if err != nil {
			log.Fatal(err)
		}
	}
}
