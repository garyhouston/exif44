package main

// Add location coordinates to a JPEG or TIFF file.

import (
	"encoding/binary"
	"errors"
	"fmt"
	exif "github.com/garyhouston/exif44"
	jseg "github.com/garyhouston/jpegsegs"
	tiff "github.com/garyhouston/tiff66"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
)

// Insert location coordinates into a GPS IFD.
func insertGPS(lat, long float64, gpsNode *tiff.IFDNode) {
	var latRef, longRef string
	if lat < 0 {
		latRef = "S"
		lat = -lat
	} else {
		latRef = "N"
	}
	if long < 0 {
		longRef = "W"
		long = -long
	} else {
		longRef = "E"
	}
	// Arbitarily setting 6 decimal places in seconds fields.
	secMult := uint32(1000000)
	latDeg := uint32(lat)
	latMin := uint32(math.Mod(lat*60, 60))
	latSec := uint32(math.Mod(lat*3600, 60) * float64(secMult))
	longDeg := uint32(long)
	longMin := uint32(math.Mod(long*60, 60))
	longSec := uint32(math.Mod(long*3600, 60) * float64(secMult))
	latData := make([]byte, tiff.RATIONAL.Size()*3)
	longData := make([]byte, tiff.RATIONAL.Size()*3)
	fields := make([]tiff.Field, 4)
	fields[0] = tiff.Field{exif.GPSLatitudeRef, tiff.ASCII, 2, []byte(latRef)}
	fields[1] = tiff.Field{exif.GPSLongitudeRef, tiff.ASCII, 2, []byte(longRef)}
	fields[2] = tiff.Field{exif.GPSLatitude, tiff.RATIONAL, 3, latData}
	fields[3] = tiff.Field{exif.GPSLongitude, tiff.RATIONAL, 3, longData}
	fields[2].PutRational(latDeg, 1, 0, gpsNode.Order)
	fields[2].PutRational(latMin, 1, 1, gpsNode.Order)
	fields[2].PutRational(latSec, secMult, 2, gpsNode.Order)
	fields[3].PutRational(longDeg, 1, 0, gpsNode.Order)
	fields[3].PutRational(longMin, 1, 1, gpsNode.Order)
	fields[3].PutRational(longSec, secMult, 2, gpsNode.Order)
	gpsNode.DeleteFields([]tiff.Tag{exif.GPSLatitudeRef, exif.GPSLatitude, exif.GPSLongitudeRef, exif.GPSLongitude})
	gpsNode.AddFields(fields)
}

// Add a GPS sub-IFD to an IFD node and return a pointer to it.
func addGPS(node *tiff.IFDNode) *tiff.IFDNode {
	gpsNode := tiff.NewIFDNode(tiff.GPSSpace)
	gpsNode.Order = node.Order
	gpsVersionData := make([]byte, 4)
	gpsVersion := tiff.Field{exif.GPSVersionID, tiff.BYTE, 4, gpsVersionData}
	gpsVersion.PutByte(2, 0)
	gpsVersion.PutByte(3, 1)
	gpsNode.AddFields([]tiff.Field{gpsVersion})
	gpsIFDData := make([]byte, 4)
	node.AddFields([]tiff.Field{{tiff.GPSIFD, tiff.LONG, 1, gpsIFDData}})
	subIFD := tiff.SubIFD{tiff.GPSIFD, gpsNode}
	node.SubIFDs = append(node.SubIFDs, subIFD)
	return gpsNode
}

// Add location coordinates into a TIFF tree, creating a GPS IFD if needed.
func putGPS(lat, long float64, root *tiff.IFDNode) {
	var gpsNode *tiff.IFDNode
	for _, i := range root.SubIFDs {
		if i.Tag == tiff.GPSIFD {
			gpsNode = i.Node
		}
	}
	if gpsNode == nil {
		gpsNode = addGPS(root)
	}
	insertGPS(lat, long, gpsNode)
}

// Process a TIFF file.
func processTIFF(lat, long float64, outfile io.Writer, infile io.Reader) error {
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
	putGPS(lat, long, tree.TIFF)
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

// Create a new Exif node for a JPEG image, from scratch.
func createExif() *exif.Exif {
	node := tiff.NewIFDNode(tiff.TIFFSpace)
	node.Order = binary.LittleEndian // arbitrary
	// Add an Exif sub-IFD too, since the Exif spec may require
	// the version to be specified.
	exifNode := tiff.NewIFDNode(tiff.ExifSpace)
	exifNode.Order = node.Order
	exifVersionData := make([]byte, 4)
	copy(exifVersionData, []byte("0230"))
	exifVersion := tiff.Field{exif.ExifVersion, tiff.UNDEFINED, 4, exifVersionData}
	exifNode.AddFields([]tiff.Field{exifVersion})
	exifIFDData := make([]byte, 4)
	node.AddFields([]tiff.Field{{tiff.ExifIFD, tiff.LONG, 1, exifIFDData}})
	subIFD := tiff.SubIFD{tiff.ExifIFD, exifNode}
	node.SubIFDs = append(node.SubIFDs, subIFD)
	return &exif.Exif{TIFF: node}
}

// Add coordinates to an Exif block and dump it to output.
func writeExif(lat, long float64, exifNode *exif.Exif, dumper *jseg.Dumper) error {
	putGPS(lat, long, exifNode.TIFF)
	app1 := make([]byte, exif.HeaderSize+exifNode.TreeSize())
	next := exif.PutHeader(app1)
	_, err := exifNode.Put(app1[next:])
	if err != nil {
		return err
	}
	return dumper.Dump(jseg.APP0+1, app1)
}

// Check if file has an Exif block.
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
		if marker == jseg.SOS {
			// Start of scan data, no more metadata expected.
			break
		}
		if marker == jseg.APP0+1 {
			haveExif, _ = exif.GetHeader(buf)
			break
		}
	}
	// Reset the file position.
	_, err = reader.Seek(readerSave, io.SeekStart)
	if err != nil {
		return false, err
	}
	return haveExif, nil
}

// Process a single image in a JPEG file. A file using Multi-Picture
// Format will contain multiple images.
func processImage(writer io.WriteSeeker, reader io.ReadSeeker, index uint32, lat, long float64, mpfProcessor jseg.MPFProcessor) error {
	var err error
	haveExif := false
	// Process Exif for the first image only.
	if index == 0 {
		// Must be done before creating the scanner.
		haveExif, err = findExif(reader)
		if err != nil {
			return err
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
	if index == 0 && !haveExif {
		// No Exif block in the file, so create one, add the GPS
		// info, and dump it to the output stream.
		exifNode := createExif()
		writeExif(lat, long, exifNode, dumper)
	}
	for {
		marker, buf, err := scanner.Scan()
		if err != nil {
			return err
		}
		if index == 0 && marker == jseg.APP0+1 {
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
				writeExif(lat, long, tree, dumper)
				continue
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
	lat        float64
	long       float64
}

// Function to be applied to each MPF image.
func (iter *iterData) MPFApply(reader io.ReadSeeker, index uint32, length uint32) error {
	if index > 0 {
		pos, err := iter.writer.Seek(0, io.SeekCurrent)
		if err != nil {
			return err
		}
		iter.newOffsets[index] = uint32(pos)
		return processImage(iter.writer, reader, index, iter.lat, iter.long, &jseg.MPFCheck{})
	}
	return nil
}

// Process additional images found in the MPF index.
func processMPFImages(writer io.WriteSeeker, reader io.ReadSeeker, lat, long float64, index *jseg.MPFIndex) ([]uint32, error) {
	var iter iterData
	iter.writer = writer
	iter.newOffsets = make([]uint32, len(index.ImageOffsets))
	iter.lat = lat
	iter.long = long
	if err := index.ImageIterate(reader, &iter); err != nil {
		return nil, err
	}
	return iter.newOffsets, nil
}

// Process a JPEG file.
func processJPEG(lat, long float64, writer io.WriteSeeker, reader io.ReadSeeker) error {
	var mpfIndex jseg.MPFIndexRewriter
	if err := processImage(writer, reader, 0, lat, long, &mpfIndex); err != nil {
		return err
	}
	if mpfIndex.Tree != nil {
		newOffsets, err := processMPFImages(writer, reader, lat, long, mpfIndex.Index)
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

func usage() {
	fmt.Printf("Usage: %s latitude longitude file outfile\nLatitude and longitude are expressed as signed decimals\n", os.Args[0])
}

// Add location coordinates to a TIFF or JPG file.
func main() {
	if len(os.Args) != 5 {
		usage()
		return
	}
	latStr := os.Args[1]
	longStr := os.Args[2]
	in := os.Args[3]
	out := os.Args[4]
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		usage()
		return
	}
	if lat < -90 || lat > 90 {
		log.Fatal("Lattitude is out of range [-90, 90]")
	}
	long, err := strconv.ParseFloat(longStr, 64)
	if err != nil {
		usage()
		return
	}
	if long < -180 || long > 180 {
		log.Fatal("Longitude is out of range [-180, 180]")
	}
	infile, err := os.Open(in)
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
	outfile, err := os.Create(out)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	if fileType == TIFFFile {
		err = processTIFF(lat, long, outfile, infile)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = processJPEG(lat, long, outfile, infile)
		if err != nil {
			log.Fatal(err)
		}
	}
}
