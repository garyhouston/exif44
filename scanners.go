package exif44

import (
	"errors"
	jseg "github.com/garyhouston/jpegsegs"
	tiff "github.com/garyhouston/tiff66"
	"io"
	"io/ioutil"
	"os"
)

// Control structure for Read and ReadFile, with optional callbacks.
type ReadControl struct {
	ReadExif ReadExif // Process Exif tree, or nil.
	// Additional callbacks could be added, e.g., for processing
	// other types of metadata, JPEG blocks, or full MPF trees.
}

type ReadExif interface {
	// Callback for processing Exif data, read-only. In the case
	// of TIFF files, this will be called once on the entire TIFF
	// tree. For JPEG files, it will be called on the Exif segment
	// for each image in the file (multiple images are supported
	// via Multi-Picture Format, MPF).
	ReadExif(imageIdx uint32, exif Exif) error
}

// Read processes its input, which is expected to be an open image
// file in a supported format, currently JPEG or TIFF. It invokes
// any callbacks in the control structure.
func Read(file io.ReadSeeker, control ReadControl) error {
	fileType, err := fileType(file)
	if err != nil {
		return err
	}
	if _, err := file.Seek(0, 0); err != nil {
		return err
	}
	if fileType == fileTIFF {
		if control.ReadExif != nil {
			if err := processTIFF(file, control); err != nil {
				return err
			}
		}
	} else {
		if err := processJPEG(file, control); err != nil {
			return err
		}
	}
	return nil
}

// ReadFile opens a file and processes it with Read.
func ReadFile(filename string, control ReadControl) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return Read(file, control)
}

// Supported file types for ReadFile and ReadWriteFile.
const (
	fileTIFF = 1
	fileJPEG = 2
)

// Determine type of stream. Anything not supported is an error. This will
// read a few bytes from the reader, changing the position.
func fileType(file io.Reader) (int, error) {
	buf := make([]byte, tiff.HeaderSize)
	if _, err := io.ReadFull(file, buf); err != nil {
		return 0, err
	}
	if jseg.IsJPEGHeader(buf) {
		return fileJPEG, nil
	}
	if validTIFF, _, _ := tiff.GetHeader(buf); validTIFF {
		return fileTIFF, nil
	}
	return 0, errors.New("File doesn't have a TIFF or JPEG header")
}

func processTIFF(file io.Reader, control ReadControl) error {
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	return scanTIFF(0, buf, control)
}

// State for the MPF image iterator.
type scanData struct {
	control ReadControl
}

// Function to be applied to each MPF image.
func (scan *scanData) MPFApply(reader io.ReadSeeker, index uint32, length uint32) error {
	if index > 0 {
		return scanJPEG(index, reader, &jseg.MPFCheck{}, scan.control)
	}
	return nil
}

func processJPEG(file io.ReadSeeker, control ReadControl) error {
	var index jseg.MPFGetIndex
	if err := scanJPEG(0, file, &index, control); err != nil {
		return err
	}
	if index.Index != nil {
		scandata := &scanData{}
		scandata.control = control
		err := index.Index.ImageIterate(file, scandata)
		if err != nil {
			return err
		}
	}
	return nil
}

// Process a single image in a JPEG file. A file using the
// Multi-Picture Format extension will contain multiple images.
func scanJPEG(imageIdx uint32, reader io.ReadSeeker, mpfProcessor jseg.MPFProcessor, control ReadControl) error {
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
		if marker == jseg.APP0+1 && control.ReadExif != nil {
			isExif, next := GetHeader(buf)
			if isExif {
				if err := scanTIFF(imageIdx, buf[next:], control); err != nil {
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

func scanTIFF(imageIdx uint32, buf []byte, control ReadControl) error {
	exif, err := GetExifTree(buf)
	if err != nil {
		return err
	}
	control.ReadExif.ReadExif(imageIdx, *exif)
	return nil
}
