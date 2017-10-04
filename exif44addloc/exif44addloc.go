package main

// Add location coordinates to a JPEG or TIFF file.

import (
	"fmt"
	exif "github.com/garyhouston/exif44"
	tiff "github.com/garyhouston/tiff66"
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

// Exif handlers.
type handlerData struct {
	latitude, longitude float64
}

func (opts handlerData) ReadWriteExif(format exif.FileFormat, imageIdx uint32, xif *exif.Exif) error {
	// Add GPS info to the first image only.
	if imageIdx == 0 {
		putGPS(opts.latitude, opts.longitude, xif.TIFF)
	}
	return nil
}

func (handlerData) ExifRequired(format exif.FileFormat, imageIdx uint32) bool {
	// Require an Exif block in the first image.
	if imageIdx == 0 {
		return true
	}
	return false
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

	var control exif.ReadWriteControl
	handlerData := handlerData{latitude: lat, longitude: long}
	control.ReadWriteExif = handlerData
	control.ExifRequired = handlerData
	if err := exif.ReadWriteFile(in, out, control); err != nil {
		log.Fatal(err)
	}
}
