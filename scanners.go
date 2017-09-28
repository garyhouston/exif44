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
func Read(reader io.ReadSeeker, control ReadControl) error {
	fileType, err := fileType(reader)
	if err != nil {
		return err
	}
	if _, err := reader.Seek(0, 0); err != nil {
		return err
	}
	if fileType == fileTIFF {
		if control.ReadExif != nil {
			if err := readTIFF(reader, control); err != nil {
				return err
			}
		}
	} else {
		if err := readJPEG(reader, control); err != nil {
			return err
		}
	}
	return nil
}

// ReadFile opens a file and processes it with Read.
func ReadFile(filename string, control ReadControl) error {
	reader, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer reader.Close()
	return Read(reader, control)
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

func readTIFF(reader io.Reader, control ReadControl) error {
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	return readTIFFBuf(0, buf, control)
}

// State for the MPF image iterator.
type scanData struct {
	control ReadControl
}

// Function to be applied to each MPF image.
func (scan *scanData) MPFApply(reader io.ReadSeeker, index uint32, length uint32) error {
	if index > 0 {
		return readJPEGImage(index, reader, &jseg.MPFCheck{}, scan.control)
	}
	return nil
}

func readJPEG(reader io.ReadSeeker, control ReadControl) error {
	var index jseg.MPFGetIndex
	if err := readJPEGImage(0, reader, &index, control); err != nil {
		return err
	}
	if index.Index != nil {
		scandata := &scanData{}
		scandata.control = control
		err := index.Index.ImageIterate(reader, scandata)
		if err != nil {
			return err
		}
	}
	return nil
}

// Process a single image in a JPEG file. A file using the
// Multi-Picture Format extension will contain multiple images.
func readJPEGImage(imageIdx uint32, reader io.ReadSeeker, mpfProcessor jseg.MPFProcessor, control ReadControl) error {
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
				if err := readTIFFBuf(imageIdx, buf[next:], control); err != nil {
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

func readTIFFBuf(imageIdx uint32, buf []byte, control ReadControl) error {
	exif, err := GetExifTree(buf)
	if err != nil {
		return err
	}
	return control.ReadExif.ReadExif(imageIdx, *exif)
}

// Control structure for ReadWrite and ReadWriteFile, with optional callbacks.
type ReadWriteControl struct {
	ReadWriteExif ReadWriteExif // Process Exif tree, or nil.
	// Additional callbacks could be added, e.g., for processing
	// other types of metadata, JPEG blocks, or full MPF trees.
}

type ReadWriteExif interface {
	// Callback for processing Exif data, read-write. In the case
	// of TIFF files, this will be called once on the entire TIFF
	// tree. For JPEG files, it will be called on the Exif segment
	// for each image in the file (multiple images are supported
	// via Multi-Picture Format, MPF). The data can be returned
	// modified or unmodified as desired.
	ReadWriteExif(imageIdx uint32, exif Exif) (Exif, error)
}

// ReadWrite processes its input, which is expected to be an open image
// file in a supported format, currently JPEG or TIFF. It invokes
// any callbacks in the control structure.
func ReadWrite(reader io.ReadSeeker, writer io.WriteSeeker, control ReadWriteControl) error {
	fileType, err := fileType(reader)
	if err != nil {
		return err
	}
	if _, err := reader.Seek(0, 0); err != nil {
		return err
	}
	if fileType == fileTIFF {
		return readWriteTIFF(reader, writer, control)
	} else {
		return readWriteJPEG(reader, writer, control)
	}
}

// ReadWriteFile opens input and output files and processes them with ReadWrite.
func ReadWriteFile(infile, outfile string, control ReadWriteControl) error {
	reader, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer reader.Close()
	writer, err := os.Create(outfile)
	if err != nil {
		return err
	}
	defer writer.Close()
	return ReadWrite(reader, writer, control)
}

func readWriteTIFF(infile io.Reader, outfile io.Writer, control ReadWriteControl) error {
	inbuf, err := ioutil.ReadAll(infile)
	if err != nil {
		return err
	}
	if inbuf == nil {
	}
	outbuf, err := readWriteTIFFBuf(0, inbuf, control)
	if err != nil {
		return err
	}
	_, err = outfile.Write(outbuf)
	return err
}

// State for MPF image iterator.
type iterData struct {
	writer     io.WriteSeeker
	newOffsets []uint32
	control    ReadWriteControl
}

// Function to be applied to each MPF image.
func (iter *iterData) MPFApply(reader io.ReadSeeker, index uint32, length uint32) error {
	if index > 0 {
		pos, err := iter.writer.Seek(0, io.SeekCurrent)
		if err != nil {
			return err
		}
		iter.newOffsets[index] = uint32(pos)
		return readWriteJPEGImage(index, reader, iter.writer, &jseg.MPFCheck{}, iter.control)
	}
	return nil
}

func readWriteJPEG(reader io.ReadSeeker, writer io.WriteSeeker, control ReadWriteControl) error {
	var mpfIndex jseg.MPFIndexRewriter
	if err := readWriteJPEGImage(0, reader, writer, &mpfIndex, control); err != nil {
		return err
	}
	if mpfIndex.Tree != nil {
		var iter iterData
		iter.writer = writer
		iter.control = control
		index := mpfIndex.Index
		iter.newOffsets = make([]uint32, len(index.ImageOffsets))
		err := index.ImageIterate(reader, &iter)
		if err != nil {
			return err
		}
		end, err := writer.Seek(0, io.SeekCurrent)
		if err != nil {
			return err
		}
		if err = jseg.RewriteMPF(writer, mpfIndex.Tree, mpfIndex.APP2WritePos, iter.newOffsets, uint32(end)); err != nil {
			return err
		}
	}
	return nil
}

// Process a single image in a JPEG file. A file using Multi-Picture
// Format will contain multiple images.
func readWriteJPEGImage(imageIdx uint32, reader io.ReadSeeker, writer io.WriteSeeker, mpfProcessor jseg.MPFProcessor, control ReadWriteControl) error {
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
			isExif, next := GetHeader(buf)
			if isExif {
				newTIFF, err := readWriteTIFFBuf(imageIdx, buf[next:], control)
				if err != nil {
					return err
				}
				buf = append(header, newTIFF...)
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

func readWriteTIFFBuf(imageIdx uint32, buf []byte, control ReadWriteControl) ([]byte, error) {
	exif, err := GetExifTree(buf)
	if err != nil {
		return nil, err
	}
	exif.TIFF.Fix()
	exifOut := *exif
	if control.ReadWriteExif != nil {
		exifOut, err = control.ReadWriteExif.ReadWriteExif(imageIdx, *exif)
		if err != nil {
			return nil, err
		}
	}
	if err = exifOut.CheckMakerNote(); err != nil {
		return nil, err
	}
	if err = exifOut.MakerNoteComplexities(); err != nil {
		return nil, err
	}
	fileSize := tiff.HeaderSize + exifOut.TreeSize()
	outbuf := make([]byte, fileSize)
	tiff.PutHeader(outbuf, exifOut.TIFF.Order, tiff.HeaderSize)
	_, err = exifOut.TIFF.PutIFDTree(outbuf, tiff.HeaderSize)
	return outbuf, err
}
