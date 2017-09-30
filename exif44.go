package exif44

import (
	"bytes"
	"errors"
	"fmt"
	tiff "github.com/garyhouston/tiff66"
)

// Tags in the Exif IFD.
const (
	ExposureTime              = 0x829A
	FNumber                   = 0x829D
	ExposureProgram           = 0x8822
	SpectralSensitivity       = 0x8824
	PhotographicSensitivity   = 0x8827
	OECF                      = 0x8828
	SensitivityType           = 0x8830
	StandardOutputSensitivity = 0x8831
	RecommendedExposureIndex  = 0x8832
	ISOSpeed                  = 0x8833
	ISOSpeedLatitudeyyy       = 0x8834
	ISOSpeedLatitudezzz       = 0x8835
	ExifVersion               = 0x9000
	DateTimeOriginal          = 0x9003
	DateTimeDigitized         = 0x9004
	ComponentsConfiguration   = 0x9101
	CompressedBitsPerPixel    = 0x9102
	ShutterSpeedValue         = 0x9201
	ApertureValue             = 0x9202
	BrightnessValue           = 0x9203
	ExposureBiasValue         = 0x9204
	MaxApertureValue          = 0x9205
	SubjectDistance           = 0x9206
	MeteringMode              = 0x9207
	LightSource               = 0x9208
	Flash                     = 0x9209
	FocalLength               = 0x920A
	SubjectArea               = 0x9214
	MakerNote                 = 0x927C
	UserComment               = 0x9286
	SubSecTime                = 0x9290
	SubSecTimeOriginal        = 0x9291
	SubSecTimeDigitized       = 0x9292
	FlashpixVersion           = 0xA000
	ColorSpace                = 0xA001
	PixelXDimension           = 0xA002
	PixelYDimension           = 0xA003
	RelatedSoundFile          = 0xA004
	InteroperabilityIFD       = 0xA005
	FlashEnergy               = 0xA20B
	SpatialFrequencyResponse  = 0xA20C
	FocalPlaneXResolution     = 0xA20E
	FocalPlaneYResolution     = 0xA20F
	FocalPlaneResolutionUnit  = 0xA210
	SubjectLocation           = 0xA214
	ExposureIndex             = 0xA215
	SensingMethod             = 0xA217
	FileSource                = 0xA300
	SceneType                 = 0xA301
	CFAPattern                = 0xA302
	CustomRendered            = 0xA401
	ExposureMode              = 0xA402
	WhiteBalance              = 0xA403
	DigitalZoomRatio          = 0xA404
	FocalLengthIn35mmFilm     = 0xA405
	SceneCaptureType          = 0xA406
	GainControl               = 0xA407
	Contrast                  = 0xA408
	Saturation                = 0xA409
	Sharpness                 = 0xA40A
	DeviceSettingDescription  = 0xA40B
	SubjectDistanceRange      = 0xA40C
	ImageUniqueID             = 0xA420
	CameraOwnerName           = 0xA430
	BodySerialNumber          = 0xA431
	LensSpecification         = 0xA432
	LensMake                  = 0xA433
	LensModel                 = 0xA434
	LensSerialNumber          = 0xA435
	Gamma                     = 0xA500
)

// Mapping from Exif tags to strings.
var ExifTagNames = map[tiff.Tag]string{
	ExposureTime:            "ExposureTime",
	FNumber:                 "FNumber",
	ExposureProgram:         "ExposureProgram",
	SpectralSensitivity:     "SpectralSensitivity",
	PhotographicSensitivity: "PhotographicSensitivity",
	OECF:                      "OECF",
	SensitivityType:           "SensitivityType",
	StandardOutputSensitivity: "StandardOutputSensitivity",
	RecommendedExposureIndex:  "RecommendedExposureIndex",
	ISOSpeed:                  "ISOSpeed",
	ISOSpeedLatitudeyyy:       "ISOSpeedLatitudeyyy",
	ISOSpeedLatitudezzz:       "ISOSpeedLatitudezzz",
	ExifVersion:               "ExifVersion",
	DateTimeOriginal:          "DateTimeOriginal",
	DateTimeDigitized:         "DateTimeDigitized",
	ComponentsConfiguration:   "ComponentsConfiguration",
	CompressedBitsPerPixel:    "CompressedBitsPerPixel",
	ShutterSpeedValue:         "ShutterSpeedValue",
	ApertureValue:             "ApertureValue",
	BrightnessValue:           "BrightnessValue",
	ExposureBiasValue:         "ExposureBiasValue",
	MaxApertureValue:          "MaxApertureValue",
	SubjectDistance:           "SubjectDistance",
	MeteringMode:              "MeteringMode",
	LightSource:               "LightSource",
	Flash:                     "Flash",
	FocalLength:               "FocalLength",
	SubjectArea:               "SubjectArea",
	MakerNote:                 "MakerNote",
	UserComment:               "UserComment",
	SubSecTime:                "SubSecTime",
	SubSecTimeOriginal:        "SubSecTimeOriginal",
	SubSecTimeDigitized:       "SubSecTimeDigitized",
	FlashpixVersion:           "FlashpixVersion",
	ColorSpace:                "ColorSpace",
	PixelXDimension:           "PixelXDimension",
	PixelYDimension:           "PixelYDimension",
	RelatedSoundFile:          "RelatedSoundFile",
	InteroperabilityIFD:       "InteroperabilityIFD",
	FlashEnergy:               "FlashEnergy",
	SpatialFrequencyResponse:  "SpatialFrequencyResponse",
	FocalPlaneXResolution:     "FocalPlaneXResolution",
	FocalPlaneYResolution:     "FocalPlaneYResolution",
	FocalPlaneResolutionUnit:  "FocalPlaneResolutionUnit",
	SubjectLocation:           "SubjectLocation",
	ExposureIndex:             "ExposureIndex",
	SensingMethod:             "SensingMethod",
	FileSource:                "FileSource",
	SceneType:                 "SceneType",
	CFAPattern:                "CFAPattern",
	CustomRendered:            "CustomRendered",
	ExposureMode:              "ExposureMode",
	WhiteBalance:              "WhiteBalance",
	DigitalZoomRatio:          "DigitalZoomRatio",
	FocalLengthIn35mmFilm:     "FocalLengthIn35mmFilm",
	SceneCaptureType:          "SceneCaptureType",
	GainControl:               "GainControl",
	Contrast:                  "Contrast",
	Saturation:                "Saturation",
	Sharpness:                 "Sharpness",
	DeviceSettingDescription:  "DeviceSettingDescription",
	SubjectDistanceRange:      "SubjectDistanceRange",
	ImageUniqueID:             "ImageUniqueID",
	CameraOwnerName:           "CameraOwnerName",
	BodySerialNumber:          "BodySerialNumber",
	LensSpecification:         "LensSpecification",
	LensMake:                  "LensMake",
	LensModel:                 "LensModel",
	LensSerialNumber:          "LensSerialNumber",
	Gamma:                     "Gamma",
}

// Tags in the Interoperability IFD, from "Design rule for Camera File
// system: DCF Version 2.0 (Edition 2010)".
const (
	InteroperabilityIndex   = 0x1
	InteroperabilityVersion = 0x2
	RelatedImageFileFormat  = 0x1000
	RelatedImageWidth       = 0x1001
	RelatedImageLength      = 0x1002
)

// Mapping from Interoperability tags to strings.
var InteropTagNames = map[tiff.Tag]string{
	InteroperabilityIndex:   "InteroperabilityIndex",
	InteroperabilityVersion: "InteroperabilityVersion",
	RelatedImageFileFormat:  "RelatedImageFileFormat",
	RelatedImageWidth:       "RelatedImageWidth",
	RelatedImageLength:      "RelatedImageLength",
}

// Tags in the GPS IFD.
const (
	GPSVersionID         = 0x00
	GPSLatitudeRef       = 0x01
	GPSLatitude          = 0x02
	GPSLongitudeRef      = 0x03
	GPSLongitude         = 0x04
	GPSAltitudeRef       = 0x05
	GPSAltitude          = 0x06
	GPSTimeStamp         = 0x07
	GPSSatellites        = 0x08
	GPSStatus            = 0x09
	GPSMeasureMode       = 0x0A
	GPSDOP               = 0x0B
	GPSSpeedRef          = 0x0C
	GPSSpeed             = 0x0D
	GPSTrackRef          = 0x0E
	GPSTrack             = 0x0F
	GPSImgDirectionRef   = 0x10
	GPSImgDirection      = 0x11
	GPSMapDatum          = 0x12
	GPSDestLatitudeRef   = 0x13
	GPSDestLatitude      = 0x14
	GPSDestLongitudeRef  = 0x15
	GPSDestLongitude     = 0x16
	GPSDestBearingRef    = 0x17
	GPSDestBearing       = 0x18
	GPSDestDistanceRef   = 0x19
	GPSDestDistance      = 0x1A
	GPSProcessingMethod  = 0x1B
	GPSAreaInformation   = 0x1C
	GPSDateStamp         = 0x1D
	GPSDifferential      = 0x1E
	GPSHPositioningError = 0x1F
)

// Mapping from GPS tags to strings.
var GPSTagNames = map[tiff.Tag]string{
	GPSVersionID:         "GPSVersionID",
	GPSLatitudeRef:       "GPSLatitudeRef",
	GPSLatitude:          "GPSLatitude",
	GPSLongitudeRef:      "GPSLongitudeRef",
	GPSLongitude:         "GPSLongitude",
	GPSAltitudeRef:       "GPSAltitudeRef",
	GPSAltitude:          "GPSAltitude",
	GPSTimeStamp:         "GPSTimeStamp",
	GPSSatellites:        "GPSSatellites",
	GPSStatus:            "GPSStatus",
	GPSMeasureMode:       "GPSMeasureMode",
	GPSDOP:               "GPSDOP",
	GPSSpeedRef:          "GPSSpeedRef",
	GPSSpeed:             "GPSSpeed",
	GPSTrackRef:          "GPSTrackRef",
	GPSTrack:             "GPSTrack",
	GPSImgDirectionRef:   "GPSImgDirectionRef",
	GPSImgDirection:      "GPSImgDirection",
	GPSMapDatum:          "GPSMapDatum",
	GPSDestLatitudeRef:   "GPSDestLatitudeRef",
	GPSDestLatitude:      "GPSDestLatitude",
	GPSDestLongitudeRef:  "GPSDestLongitudeRef",
	GPSDestLongitude:     "GPSDestLongitude",
	GPSDestBearingRef:    "GPSDestBearingRef",
	GPSDestBearing:       "GPSDestBearing",
	GPSDestDistanceRef:   "GPSDestDistanceRef",
	GPSDestDistance:      "GPSDestDistance",
	GPSProcessingMethod:  "GPSProcessingMethod",
	GPSAreaInformation:   "GPSAreaInformation",
	GPSDateStamp:         "GPSDateStamp",
	GPSDifferential:      "GPSDifferential",
	GPSHPositioningError: "GPSHPositioningError",
}

// Return the tag->name map for given namespace, or a nil map if the space
// is unknown.
func TagNameMap(space tiff.TagSpace) map[tiff.Tag]string {
	var names map[tiff.Tag]string
	switch space {
	case tiff.TIFFSpace:
		names = tiff.TagNames
	case tiff.ExifSpace:
		names = ExifTagNames
	case tiff.InteropSpace:
		names = InteropTagNames
	case tiff.GPSSpace:
		names = GPSTagNames
	case tiff.Canon1Space:
		names = Canon1TagNames
	case tiff.Olympus1Space:
		names = Olympus1TagNames
	case tiff.Olympus1EquipmentSpace:
		names = Olympus1EquipmentTagNames
	case tiff.Olympus1CameraSettingsSpace:
		names = Olympus1CameraSettingsTagNames
	case tiff.Olympus1RawDevelopmentSpace:
		names = Olympus1RawDevelopmentTagNames
	case tiff.Olympus1RawDev2Space:
		names = Olympus1RawDev2TagNames
	case tiff.Olympus1ImageProcessingSpace:
		names = Olympus1ImageProcessingTagNames
	case tiff.Olympus1FocusInfoSpace:
		names = Olympus1FocusInfoTagNames
	case tiff.Panasonic1Space:
		names = Panasonic1TagNames
	case tiff.Nikon1Space:
		names = Nikon1TagNames
	case tiff.Nikon2Space:
		names = Nikon2TagNames
	case tiff.Nikon2PreviewSpace:
		names = Nikon2PreviewIFDTagNames
	case tiff.Nikon2ScanSpace:
		names = Nikon2ScanIFDTagNames
	}
	return names
}

// Exif header, as found in a JPEG APP1 segment.
var header = []byte("Exif\000\000")

// Size of an Exif header.
const HeaderSize = 6

// Check if a slice starts with an Exif header, as found in a JPEG
// APP1 segment.  Returns a flag and the position of the next byte.
func GetHeader(buf []byte) (bool, uint32) {
	if uint32(len(buf)) >= HeaderSize && bytes.Compare(buf[:HeaderSize], header) == 0 {
		return true, HeaderSize
	} else {
		return false, 0
	}
}

// Put an Exif header, as for a JPEG APP1 segment, at the start of a slice,
// returning the position of the next byte.
func PutHeader(buf []byte) uint32 {
	copy(buf, header)
	return HeaderSize
}

type Exif struct {
	TIFF      *tiff.IFDNode // Pointer to top-level TIFF node, IFD0. This is the root of the tree containing all other nodes.
	Exif      *tiff.IFDNode // Pointer to Exif in tree.
	GPS       *tiff.IFDNode // Pointer to GPS in tree.
	Interop   *tiff.IFDNode // Pointer to Interop in tree.
	MakerNote *tiff.IFDNode // Pointer to maker note in tree.
}

// Create an Exif struct of pointers from an IFD tree.
func makeExif(node *tiff.IFDNode) *Exif {
	exif := Exif{TIFF: node}
	for _, sub := range node.SubIFDs {
		if sub.Node.GetSpace() == tiff.ExifSpace {
			exif.Exif = sub.Node
			for _, esub := range sub.Node.SubIFDs {
				if esub.Node.GetSpace() == tiff.InteropSpace {
					exif.Interop = esub.Node
				} else if esub.Node.IsMakerNote() {
					exif.MakerNote = esub.Node
				}
			}
		} else if sub.Node.GetSpace() == tiff.GPSSpace {
			exif.GPS = sub.Node
		}
	}
	return &exif
}

// Unpack a TIFF header and tree from a slice, using GetHeader and
// GetIFDTree from tiff66, and also return pointers to any Exif
// subIFDs that are present. As for GetIFDTree, field data in the
// structures points into the original byte slice, so modifying one
// will modify the other.
func GetExifTree(buf []byte) (*Exif, error) {
	valid, order, ifdpos := tiff.GetHeader(buf)
	if !valid {
		return nil, errors.New("GetExifTree: Invalid Tiff header")
	}
	node, err := tiff.GetIFDTree(buf, order, ifdpos, tiff.TIFFSpace)
	if err != nil {
		return nil, err
	}
	return makeExif(node), nil
}

// Return the size of a buffer needed to serialize the tree in an Exif
// structure in TIFF format, including the TIFF header, but excluding
// the Exif header used in JPEG files.
func (exif Exif) TreeSize() uint32 {
	return tiff.HeaderSize + exif.TIFF.TreeSize()
}

// Pack Exif data into a slice in TIFF format. The slice should start
// with the first byte following any Exif header. Returns the position
// following the last byte used.
func (exif *Exif) Put(buf []byte) (uint32, error) {
	tiff.PutHeader(buf, exif.TIFF.Order, tiff.HeaderSize)
	return exif.TIFF.PutIFDTree(buf, tiff.HeaderSize)
}

func allZero(s []byte) bool {
	for i := range s {
		if s[i] != 0 {
			return false
		}
	}
	return true
}

// Return an error if an Exif tree contains a maker note that wasn't
// decoded, and may be damaged if rewritten at a new location.
func (exif Exif) CheckMakerNote() error {
	if exif.Exif != nil {
		fields := exif.Exif.FindFields([]tiff.Tag{MakerNote})
		if len(fields) > 0 && exif.MakerNote == nil {
			maker := fields[0].Data
			// Panasonic PV-DC2090 (c1999) creates maker notes with
			// four zero-valued bytes. Ignore all-zero maker
			// notes, since they won't be damaged by relocation.
			if allZero(maker) {
				return nil
			}
			plen := len(maker)
			cont := ""
			if plen > 15 {
				plen = 15
				cont = "..."
			}
			return errors.New(fmt.Sprintf("Unsupported maker note: %q%s", maker[0:plen], cont))
		}
	}
	return nil
}

// Return an error if an Exif tree contains a maker note that requires
// special processing (i.e., library functionality to get and put the
// Exif tree won't process it correctly.) This is for applications
// that don't bother to handle these cases.
func (exif Exif) MakerNoteComplexities() error {
	if exif.MakerNote != nil {
		if exif.MakerNote.GetSpace() == tiff.Canon1Space {
			// Preview images in Canon EOS 300D maker notes
			// are stored outside the Exif JPEG segment.
			fields := exif.MakerNote.FindFields([]tiff.Tag{Canon1PreviewImageInfo})
			if len(fields) > 0 {
				return errors.New(fmt.Sprintf("Unsupported PreviewImageInfo field in Canon maker note"))
			}
		}
	}
	return nil
}
