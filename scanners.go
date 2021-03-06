package exif44

import (
	"encoding/binary"
	"errors"
	jseg "github.com/garyhouston/jpegsegs"
	tiff "github.com/garyhouston/tiff66"
	"io"
	"io/ioutil"
	"os"
)

// Control structure for Read and ReadFile, with optional callbacks.
type ReadControl struct {
	ReadExif ReadExif // Callback to process Exif tree, or nil.
	// Additional callbacks could be added, e.g., for processing
	// other types of metadata, JPEG blocks, or full MPF trees.
}

type ReadExif interface {
	// Callback for processing Exif data, read-only. In the case
	// of TIFF files, it will be called once on each image in the
	// TIFF tree, which are linked together using the Next
	// pointers.  For JPEG files, it will be called on the Exif
	// segment for each image in the file (multiple images are
	// supported via Multi-Picture Format, MPF), and the Next
	// pointer may link to a thumbnail image. Any errors from
	// decoding the data will be available in err, which may be
	// a multierror structure. Returning a non-nil error will
	// terminate processing.
	ReadExif(format FileFormat, imageIdx uint32, exif Exif, err error) error
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
	if fileType == FileTIFF {
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

// Supported file formats for ReadFile and ReadWriteFile.
type FileFormat uint8

const (
	FileTIFF = 1
	FileJPEG = 2
)

// Determine type of stream. Anything not supported is an error. This will
// read a few bytes from the reader, changing the position.
func fileType(file io.Reader) (FileFormat, error) {
	buf := make([]byte, tiff.HeaderSize)
	if _, err := io.ReadFull(file, buf); err != nil {
		return 0, err
	}
	if jseg.IsJPEGHeader(buf) {
		return FileJPEG, nil
	}
	if validTIFF, _, _ := tiff.GetHeader(buf); validTIFF {
		return FileTIFF, nil
	}
	return 0, errors.New("File doesn't have a TIFF or JPEG header")
}

func readTIFF(reader io.Reader, control ReadControl) error {
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	return readTIFFBuf(FileTIFF, 0, buf, control)
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
		if marker == jseg.SOS || marker == jseg.EOI {
			// No more metadata expected.
			return nil
		}
		if marker == jseg.APP0+1 && control.ReadExif != nil {
			isExif, next := GetHeader(buf)
			if isExif {
				// Copy the buffer so that data in the Exif tree can remain valid if the callback decides to save it.
				copyBuf := make([]byte, len(buf[next:]))
				copy(copyBuf, buf[next:])
				if err := readTIFFBuf(FileJPEG, imageIdx, copyBuf, control); err != nil {
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

func readTIFFBuf(format FileFormat, imageIdx uint32, buf []byte, control ReadControl) error {
	exif, err := GetExifTree(buf)
	for {
		if err = control.ReadExif.ReadExif(format, imageIdx, *exif, err); err != nil {
			return err
		}
		if format == FileJPEG || exif.TIFF.Next == nil {
			return nil
		}
		exif = makeExif(exif.TIFF.Next)
		imageIdx++
	}
}

// Control structure for ReadWrite and ReadWriteFile, with optional callbacks.
type ReadWriteControl struct {
	ReadWriteExif ReadWriteExif // Callback to process Exif tree, or nil.
	ExifRequired  ExifRequired  // Check whether Exif block should be added if not present.

	// Additional callbacks could be added, e.g., for processing
	// other types of metadata, JPEG blocks, or full MPF trees.
}

type ReadWriteExif interface {
	// Callback for processing Exif data, read-write. In the case
	// of TIFF files, it will be called once on each image in the
	// TIFF tree, which are linked together using the Next
	// pointers.  For JPEG files, it will be called on the Exif
	// segment for each image in the file (multiple images are
	// supported via Multi-Picture Format, MPF), and the Next
	// pointer may link to a thumbnail image. Any errors from
	// decoding the data will be available in err, which may be
	// a multierror structure. Returning a non-nil error will
	// terminate processing.
	ReadWriteExif(format FileFormat, imageIdx uint32, exif *Exif, err error) error
}

type ExifRequired interface {
	// Callback to determine whether an Exif block should be
	// created if not already present for the specfied image
	// number. For a JPEG file, an APP1 segment will be created if
	// necessary. For JPEG or TIFF, an Exif IFD will be created
	// containing an ExifVersion field.
	ExifRequired(format FileFormat, imageIdx uint32) bool
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
	if fileType == FileTIFF {
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
	outbuf, err := readWriteTIFFBuf(FileTIFF, 0, inbuf, control)
	if outbuf == nil && err == nil {
		err = errors.New("TIFF file would contain no fields, not writing.")
	}
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
	needExif := control.ExifRequired != nil && control.ExifRequired.ExifRequired(FileJPEG, imageIdx)
	if needExif {
		// Must be done before creating the scanner.
		haveExif, err := findExif(reader)
		if err != nil {
			return err
		}
		if haveExif {
			needExif = false
		}
	}
	scanner, err := jseg.NewScanner(reader)
	if err != nil {
		return err
	}
	dumper, err := jseg.NewDumper(writer)
	if err != nil {
		return err
	}
	if needExif {
		// Create an Exif segment at the start of the output.
		buf, err := createExif(imageIdx, control)
		if err != nil {
			return err
		}
		if buf != nil {
			if err = dumper.Dump(jseg.APP0+1, buf); err != nil {
				return err
			}
		}
	}
	for {
		marker, buf, err := scanner.Scan()
		if err != nil {
			return err
		}
		if marker == jseg.APP0+1 {
			isExif, next := GetHeader(buf)
			if isExif {
				// Copy the buffer so that data in the Exif tree can remain valid if the callback decides to save it.
				copyBuf := make([]byte, len(buf[next:]))
				copy(copyBuf, buf[next:])
				newTIFF, err := readWriteTIFFBuf(FileJPEG, imageIdx, copyBuf, control)
				if err != nil {
					return err
				}
				if newTIFF == nil {
					// Skip empty TIFF tree.
					continue
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

// Check if a JPEG image has an Exif block.
func findExif(reader io.ReadSeeker) (bool, error) {
	readerSave, err := reader.Seek(0, io.SeekCurrent)
	if err != nil {
		return false, err
	}
	scanner, err := jseg.NewScanner(reader)
	if err != nil {
		return false, err
	}
	haveExif := false
	for {
		marker, buf, err := scanner.Scan()
		if err != nil {
			return false, err
		}
		if marker == jseg.SOS || marker == jseg.EOI {
			// No more metadata expected.
			break
		}
		if marker == jseg.APP0+1 {
			haveExif, _ = GetHeader(buf)
			if haveExif {
				break
			}
		}
	}
	// Reset the file position.
	_, err = reader.Seek(readerSave, io.SeekStart)
	if err != nil {
		return false, err
	}
	return haveExif, nil
}

// Create an empty Exif node, call the ReadWrite callback on it, and
// serialize the result. Returns nil if the callback empties the node.
func createExif(imageIdx uint32, control ReadWriteControl) ([]byte, error) {
	node := tiff.NewIFDNode(tiff.TIFFSpace)
	node.Order = binary.LittleEndian // arbitrary
	exif := Exif{TIFF: node}
	bufSize := tiff.HeaderSize + exif.TreeSize()
	buf := make([]byte, bufSize)
	tiff.PutHeader(buf, exif.TIFF.Order, tiff.HeaderSize)
	if _, err := exif.TIFF.PutIFDTree(buf, tiff.HeaderSize); err != nil {
		return nil, err
	}
	newTIFF, err := readWriteTIFFBuf(FileJPEG, imageIdx, buf, control)
	if newTIFF == nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	buf = append(header, newTIFF...)
	return buf, nil
}

// Create an Exif IFD and add it to a TIFF tree.
func addExifIFD(exif *Exif) {
	// Create the Exif IFD node.
	exifNode := tiff.NewIFDNode(tiff.ExifSpace)
	exifNode.Order = exif.TIFF.Order
	// Add the version field to the node.
	exifVersionData := make([]byte, 4)
	copy(exifVersionData, []byte("0230"))
	exifVersion := tiff.Field{Tag: ExifVersion, Type: tiff.UNDEFINED, Count: 4, Data: exifVersionData}
	exifNode.AddFields([]tiff.Field{exifVersion})
	// Add a ExifIFD field to the TIFF IFD. Data will be set to the right
	// offset when the tree is serialized.
	exifIFDData := make([]byte, 4)
	tiffNode := exif.TIFF
	tiffNode.AddFields([]tiff.Field{{Tag: tiff.ExifIFD, Type: tiff.LONG, Count: 1, Data: exifIFDData}})
	// Add the Exif node to the TIFF node's sub-IFD list.
	subIFD := tiff.SubIFD{tiff.ExifIFD, exifNode}
	tiffNode.SubIFDs = append(tiffNode.SubIFDs, subIFD)
	// Set the pointer in the Exif struct.
	exif.Exif = exifNode
}

// Given a tiff buffer, applies callbacks and returns a newly
// allocated buffer, or nil if an error occurs or if there was no
// output to be written.
func readWriteTIFFBuf(format FileFormat, imageIdx uint32, buf []byte, control ReadWriteControl) ([]byte, error) {
	exif, err := GetExifTree(buf)
	exif.TIFF.Fix()
	exifNode := exif
	for exifNode != nil {
		if exifNode.Exif == nil && control.ExifRequired != nil && control.ExifRequired.ExifRequired(format, imageIdx) == true {
			addExifIFD(exifNode)
		}
		if control.ReadWriteExif != nil {
			if err = control.ReadWriteExif.ReadWriteExif(format, imageIdx, exifNode, err); err != nil {
				return nil, err
			}
		}
		if err = exifNode.CheckMakerNote(); err != nil {
			return nil, err
		}
		if err = exifNode.MakerNoteComplexities(); err != nil {
			return nil, err
		}
		if format == FileJPEG || exifNode.TIFF.Next == nil {
			exifNode = nil
		} else {
			exifNode = makeExif(exifNode.TIFF.Next)
			imageIdx++
		}
	}
	exif.TIFF = exif.TIFF.DeleteEmptyIFDs()
	if exif.TIFF == nil {
		// No fields, empty output.
		return nil, nil
	}
	bufSize := tiff.HeaderSize + exif.TreeSize()
	outbuf := make([]byte, bufSize)
	tiff.PutHeader(outbuf, exif.TIFF.Order, tiff.HeaderSize)
	_, err = exif.TIFF.PutIFDTree(outbuf, tiff.HeaderSize)
	return outbuf, err
}
