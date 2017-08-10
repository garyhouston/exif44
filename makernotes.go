package exif44

import (
	tiff "github.com/garyhouston/tiff66"
)

// Tag names are from ExifTool 10.56, which also has information about
// how to interpret the values.

// Tags in the Canon1 maker note, used for all Canon models.
const (
	Canon1CameraSettings             = 0x1
	Canon1FocalLength                = 0x2
	Canon1FlashInfo                  = 0x3
	Canon1ShotInfo                   = 0x4
	Canon1Panorama                   = 0x5
	Canon1ImageType                  = 0x6
	Canon1FirmwareVersion            = 0x7
	Canon1FileNumber                 = 0x8
	Canon1OwnerName                  = 0x9
	Canon1SerialNumber               = 0xc
	Canon1CameraInfo                 = 0xd
	Canon1FileLength                 = 0xe
	Canon1CustomFunctions            = 0xf
	Canon1ModelID                    = 0x10
	Canon1MovieInfo                  = 0x11
	Canon1AFInfo                     = 0x12
	Canon1ThumnailImageValidArea     = 0x13
	Canon1SerialNumberFormat         = 0x15
	Canon1SuperMacro                 = 0x1a
	Canon1DateStampMode              = 0x1c
	Canon1MyColors                   = 0x1d
	Canon1FirmwareRevision           = 0x1e
	Canon1Categories                 = 0x23
	Canon1FaceDetect1                = 0x24
	Canon1FaceDetect2                = 0x25
	Canon1AFInfo2                    = 0x26
	Canon1ContrastInfo               = 0x27
	Canon1ImageUniqueID              = 0x28
	Canon1FaceDetect3                = 0x2f
	Canon1TimeInfo                   = 0x35
	Canon1AFInfo3                    = 0x3c
	Canon1RawDataOffset              = 0x81
	Canon1OriginalDecisionDataOffset = 0x83
	Canon1CustomFunctions1D          = 0x90
	Canon1PersonalFunctions          = 0x91
	Canon1PersonalFunctionValues     = 0x92
	Canon1FileInfo                   = 0x93
	Canon1AFPointsInFocus1D          = 0x94
	Canon1LensModel                  = 0x95
	Canon1InternalSerialNumber       = 0x96
	Canon1DustRemovalData            = 0x97
	Canon1CropInfo                   = 0x98
	Canon1CustomFunctions2           = 0x99
	Canon1AspectInfo                 = 0x9a
	Canon1ProcessingInfo             = 0xa0
	Canon1ToneCurveTable             = 0xa1
	Canon1SharpnessTable             = 0xa2
	Canon1SharpnessFreqTable         = 0xa3
	Canon1WhiteBalanceTable          = 0xa4
	Canon1ColorBalance               = 0xa9
	Canon1MeasuredColor              = 0xaa
	Canon1ColorTemperature           = 0xae
	Canon1Flags                      = 0xb0
	Canon1ModifiedInfo               = 0xb1
	Canon1ToneCurveMatching          = 0xb2
	Canon1WhiteBalanceMatching       = 0xb3
	Canon1ColorSpace                 = 0xb4
	Canon1PreviewImageInfo           = 0xb6
	Canon1VRDOffset                  = 0xd0
	Canon1SensorInfo                 = 0xe0
	Canon1ColorData                  = 0x4001
	Canon1CRWParam                   = 0x4002
	Canon1ColorInfo                  = 0x4003
	Canon1Flavor                     = 0x4005
	Canon1PictureStyleUserDef        = 0x4008
	Canon1PictureStylePC             = 0x4009
	Canon1CustomPictureStyleFileName = 0x4010
	Canon1AFMicroAdj                 = 0x4013
	Canon1VignettingCorr             = 0x4015
	Canon1VignettingCorr2            = 0x4016
	Canon1LightingOpt                = 0x4018
	Canon1LensInfo                   = 0x4019
	Canon1AmbienceInfo               = 0x4020
	Canon1MultiExp                   = 0x4021
	Canon1FilterInfo                 = 0x4024
	Canon1HDRInfo                    = 0x4025
	Canon1AFConfig                   = 0x4028
)

// Mapping from Canon1 tags to strings.
var Canon1TagNames = map[tiff.Tag]string{
	Canon1CameraSettings:             "CameraSettings",
	Canon1FocalLength:                "FocalLength",
	Canon1FlashInfo:                  "FlashInfo",
	Canon1ShotInfo:                   "ShotInfo",
	Canon1Panorama:                   "Panorama",
	Canon1ImageType:                  "ImageType",
	Canon1FirmwareVersion:            "FirmwareVersion",
	Canon1FileNumber:                 "FileNumber",
	Canon1OwnerName:                  "OwnerName",
	Canon1SerialNumber:               "SerialNumber",
	Canon1CameraInfo:                 "CameraInfo",
	Canon1FileLength:                 "FileLength",
	Canon1CustomFunctions:            "CustomFunctions",
	Canon1ModelID:                    "ModelID",
	Canon1MovieInfo:                  "MovieInfo",
	Canon1AFInfo:                     "AFInfo",
	Canon1ThumnailImageValidArea:     "ThumnailImageValidArea",
	Canon1SerialNumberFormat:         "SerialNumberFormat",
	Canon1SuperMacro:                 "SuperMacro",
	Canon1DateStampMode:              "DateStampMode",
	Canon1MyColors:                   "MyColors",
	Canon1FirmwareRevision:           "FirmwareRevision",
	Canon1Categories:                 "Categories",
	Canon1FaceDetect1:                "FaceDetect1",
	Canon1FaceDetect2:                "FaceDetect2",
	Canon1AFInfo2:                    "AFInfo2",
	Canon1ContrastInfo:               "ContrastInfo",
	Canon1ImageUniqueID:              "ImageUniqueID",
	Canon1FaceDetect3:                "FaceDetect3",
	Canon1TimeInfo:                   "TimeInfo",
	Canon1AFInfo3:                    "AFInfo3",
	Canon1RawDataOffset:              "RawDataOffset",
	Canon1OriginalDecisionDataOffset: "OriginalDecisionDataOffset",
	Canon1CustomFunctions1D:          "CustomFunctions1D",
	Canon1PersonalFunctions:          "PersonalFunctions",
	Canon1PersonalFunctionValues:     "PersonalFunctionValues",
	Canon1FileInfo:                   "FileInfo",
	Canon1AFPointsInFocus1D:          "AFPointsInFocus1D",
	Canon1LensModel:                  "LensModel",
	Canon1InternalSerialNumber:       "InternalSerialNumber",
	Canon1DustRemovalData:            "DustRemovalData",
	Canon1CropInfo:                   "CropInfo",
	Canon1CustomFunctions2:           "CustomFunctions2",
	Canon1AspectInfo:                 "AspectInfo",
	Canon1ProcessingInfo:             "ProcessingInfo",
	Canon1ToneCurveTable:             "ToneCurveTable",
	Canon1SharpnessTable:             "SharpnessTable",
	Canon1SharpnessFreqTable:         "SharpnessFreqTable",
	Canon1WhiteBalanceTable:          "WhiteBalanceTable",
	Canon1ColorBalance:               "ColorBalance",
	Canon1MeasuredColor:              "MeasuredColor",
	Canon1ColorTemperature:           "ColorTemperature",
	Canon1Flags:                      "Flags",
	Canon1ModifiedInfo:               "ModifiedInfo",
	Canon1ToneCurveMatching:          "ToneCurveMatching",
	Canon1WhiteBalanceMatching:       "WhiteBalanceMatching",
	Canon1ColorSpace:                 "ColorSpace",
	Canon1PreviewImageInfo:           "PreviewImageInfo",
	Canon1VRDOffset:                  "VRDOffset",
	Canon1SensorInfo:                 "SensorInfo",
	Canon1ColorData:                  "ColorData",
	Canon1CRWParam:                   "CRWParam",
	Canon1ColorInfo:                  "ColorInfo",
	Canon1Flavor:                     "Flavor",
	Canon1PictureStyleUserDef:        "PictureStyleUserDef",
	Canon1PictureStylePC:             "PictureStylePC",
	Canon1CustomPictureStyleFileName: "CustomPictureStyleFileName",
	Canon1AFMicroAdj:                 "AFMicroAdj",
	Canon1VignettingCorr:             "VignettingCorr",
	Canon1VignettingCorr2:            "VignettingCorr2",
	Canon1LightingOpt:                "LightingOpt",
	Canon1LensInfo:                   "LensInfo",
	Canon1AmbienceInfo:               "AmbienceInfo",
	Canon1MultiExp:                   "MultiExp",
	Canon1FilterInfo:                 "FilterInfo",
	Canon1HDRInfo:                    "HDRInfo",
	Canon1AFConfig:                   "AFConfig",
}

// Tags in the Nikon1 maker note, used in early digital cameras
// such as Coolpix 700, 800, 900 and 950. Exiftool calls it Type 2,
// other sites call it Type 1.  Exiftool: lib/Image/ExifTool/Nikon.pm,
// Image::ExifTool::Nikon::Type2.
const (
	Nikon1Quality         = 0x3
	Nikon1ColorMode       = 0x4
	Nikon1ImageAdjustment = 0x5
	Nikon1CCDSensitivity  = 0x6
	Nikon1WhiteBalance    = 0x7
	Nikon1Focus           = 0x8
	Nikon1DigitalZoom     = 0xA
	Nikon1Converter       = 0xB
)

// Mapping from Nikon1 tags to strings.
var Nikon1TagNames = map[tiff.Tag]string{
	Nikon1Quality:         "Quality",
	Nikon1ColorMode:       "ColorMode",
	Nikon1ImageAdjustment: "ImageAdjustment",
	Nikon1CCDSensitivity:  "CCDSensitivity",
	Nikon1WhiteBalance:    "WhiteBalance",
	Nikon1Focus:           "Focus",
	Nikon1DigitalZoom:     "DigitalZoom",
	Nikon1Converter:       "Converter",
}

// Tags that are be used by Nikon since 2000 or so, such as in the
// Coolpix 5000.  The maker note header may be missing for early
// cameras, and has several versions otherwise.  Exiftool:
// lib/Image/ExifTool/Nikon.pm, Image::ExifTool::Nikon::Main.

const (
	Nikon2MakerNoteVersion          = 0x1
	Nikon2ISO                       = 0x2
	Nikon2ColorMode                 = 0x3
	Nikon2Quality                   = 0x4
	Nikon2WhiteBalance              = 0x5
	Nikon2Sharpness                 = 0x6
	Nikon2FocusMode                 = 0x7
	Nikon2FlashSetting              = 0x8
	Nikon2FlashType                 = 0x9
	Nikon2WhiteBalanceFineTune      = 0xB
	Nikon2WBRBLevels                = 0XC
	Nikon2ProgramShift              = 0XD
	Nikon2ExposureDifference        = 0XE
	Nikon2ISOSelection              = 0xF
	Nikon2DataDump                  = 0x10
	Nikon2PreviewIFD                = 0x11
	Nikon2FlashExposureComp         = 0x12
	Nikon2ISOSetting                = 0x13
	Nikon2ColorBalanceA             = 0x14
	Nikon2ImageBoundary             = 0x16
	Nikon2ExternalFlashExposureComp = 0x17
	Nikon2FlashExposureBracketVlaue = 0x18
	Nikon2ExposureBracketValue      = 0x19
	Nikon2ImageProcessing           = 0x1A
	Nikon2CropHiSpeed               = 0x1B
	Nikon2ExposureTuning            = 0x1C
	Nikon2SerialNumber              = 0x1D
	Nikon2ColorSpace                = 0x1E
	Nikon2VRInfo                    = 0x1F
	Nikon2ImageAuthentication       = 0x20
	Nikon2FaceDetect                = 0x21
	Nikon2ActiveDLighting           = 0x22
	Nikon2PictureControlData        = 0x23
	Nikon2WorldTime                 = 0x24
	Nikon2ISOInfo                   = 0x25
	Nikon2VignetteControl           = 0x2A
	Nikon2DistortInfo               = 0x2B
	Nikon2HDRInfo                   = 0x35
	Nikon2LocationInfo              = 0x39
	Nikon2BlackLevel                = 0x3D
	Nikon2ImageAdjustment           = 0x80
	Nikon2ToneComp                  = 0x81
	Nikon2AuxiliaryLens             = 0x82
	Nikon2LensType                  = 0x83
	Nikon2Lens                      = 0x84
	Nikon2ManualFocusDistance       = 0x85
	Nikon2DigitalZoom               = 0x86
	Nikon2FlashMode                 = 0x87
	Nikon2AFInfo                    = 0x88
	Nikon2ShootingMode              = 0x89
	Nikon2LensFStops                = 0x8B
	Nikon2ConstrastCurve            = 0x8C
	Nikon2ColorHue                  = 0x8D
	Nikon2SceneMode                 = 0x8F
	Nikon2LightSource               = 0x90
	Nikon2ShotInfo                  = 0x91
	Nikon2HueAdjustment             = 0x92
	Nikon2NEFCompression            = 0x93
	Nikon2Saturation                = 0x94
	Nikon2NoiseReduction            = 0x95
	Nikon2NEFLinearizationTable     = 0x96
	Nikon2ColorBalance              = 0x97
	Nikon2LensData                  = 0x98
	Nikon2RawImageCentre            = 0x99
	Nikon2SensorPixelSize           = 0x9A
	Nikon2SceneAssist               = 0x9C
	Nikon2RetouchHistory            = 0x9E
	Nikon2SerialNumber2             = 0xA0
	Nikon2ImageDataSize             = 0xA2
	Nikon2ImageCount                = 0xA5
	Nikon2DeletedImageCount         = 0xA6
	Nikon2ShutterCount              = 0xA7
	Nikon2FlashInfo                 = 0xA8
	Nikon2ImageOptimization         = 0xA9
	Nikon2Saturation2               = 0xAA
	Nikon2VariProgram               = 0xAB
	Nikon2ImageStabilization        = 0xAC
	Nikon2AFResponse                = 0xAD
	Nikon2MultiExposure             = 0xB0
	Nikon2HighISONoiseReduction     = 0xB1
	Nikon2ToningEffect              = 0xB3
	Nikon2PowerUpTime               = 0xB6
	Nikon2AFInfo2                   = 0xB7
	Nikon2FileInfo                  = 0xB8
	Nikon2AFTune                    = 0xB9
	Nikon2RetouchInfo               = 0xBB
	Nikon2PictureControlData2       = 0xBD
	Nikon2BarometerInfo             = 0xC3
	Nikon2PrintIM                   = 0xE00
	Nikon2NikonCaptureData          = 0xE01
	Nikon2NikonCaptureVersion       = 0xE09
	Nikon2NikonCaptureOffsets       = 0xE0E
	Nikon2NikonScanIFD              = 0xE10
	Nikon2NikonCaptureEditVersions  = 0xE13
	Nikon2NikonICCProfile           = 0xE1D
	Nikon2NikonCaptureOutput        = 0xE1E
	Nikon2NEFBitDepth               = 0xE22
)

// Mapping from Nikon2 tags to strings.
var Nikon2TagNames = map[tiff.Tag]string{
	Nikon2MakerNoteVersion:          "MakerNoteVersion",
	Nikon2ISO:                       "ISO",
	Nikon2ColorMode:                 "ColorMode",
	Nikon2Quality:                   "Quality",
	Nikon2WhiteBalance:              "WhiteBalance",
	Nikon2Sharpness:                 "Sharpness",
	Nikon2FocusMode:                 "FocusMode",
	Nikon2FlashSetting:              "FlashSetting",
	Nikon2FlashType:                 "FlashType",
	Nikon2WhiteBalanceFineTune:      "WhiteBalanceFineTune",
	Nikon2WBRBLevels:                "WB_RBLevels",
	Nikon2ProgramShift:              "ProgramShift",
	Nikon2ExposureDifference:        "ExposureDifference",
	Nikon2ISOSelection:              "ISOSelection",
	Nikon2DataDump:                  "DataDump",
	Nikon2PreviewIFD:                "PreviewIFD",
	Nikon2FlashExposureComp:         "FlashExposureComp",
	Nikon2ISOSetting:                "ISOSetting",
	Nikon2ColorBalanceA:             "ColorBalanceA",
	Nikon2ImageBoundary:             "ImageBoundary",
	Nikon2ExternalFlashExposureComp: "ExternalFlashExposureComp",
	Nikon2FlashExposureBracketVlaue: "FlashExposureBracketVlaue",
	Nikon2ExposureBracketValue:      "ExposureBracketValue",
	Nikon2ImageProcessing:           "ImageProcessing",
	Nikon2CropHiSpeed:               "CropHiSpeed",
	Nikon2ExposureTuning:            "ExposureTuning",
	Nikon2SerialNumber:              "SerialNumber",
	Nikon2ColorSpace:                "ColorSpace",
	Nikon2VRInfo:                    "VRInfo",
	Nikon2ImageAuthentication:       "ImageAuthentication",
	Nikon2FaceDetect:                "FaceDetect",
	Nikon2ActiveDLighting:           "ActiveD-Lighting",
	Nikon2PictureControlData:        "PictureControlData",
	Nikon2WorldTime:                 "WorldTime",
	Nikon2ISOInfo:                   "ISOInfo",
	Nikon2VignetteControl:           "VignetteControl",
	Nikon2DistortInfo:               "DistortInfo",
	Nikon2HDRInfo:                   "HDRInfo",
	Nikon2LocationInfo:              "LocationInfo",
	Nikon2BlackLevel:                "BlackLevel",
	Nikon2ImageAdjustment:           "ImageAdjustment",
	Nikon2ToneComp:                  "ToneComp",
	Nikon2AuxiliaryLens:             "AuxiliaryLens",
	Nikon2LensType:                  "LensType",
	Nikon2Lens:                      "Lens",
	Nikon2ManualFocusDistance:       "ManualFocusDistance",
	Nikon2DigitalZoom:               "DigitalZoom",
	Nikon2FlashMode:                 "FlashMode",
	Nikon2AFInfo:                    "AFInfo",
	Nikon2ShootingMode:              "ShootingMode",
	Nikon2LensFStops:                "LensFStops",
	Nikon2ConstrastCurve:            "ConstrastCurve",
	Nikon2ColorHue:                  "ColorHue",
	Nikon2SceneMode:                 "SceneMode",
	Nikon2LightSource:               "LightSource",
	Nikon2ShotInfo:                  "ShotInfo",
	Nikon2HueAdjustment:             "HueAdjustment",
	Nikon2NEFCompression:            "NEFCompression",
	Nikon2Saturation:                "Saturation",
	Nikon2NoiseReduction:            "NoiseReduction",
	Nikon2NEFLinearizationTable:     "NEFLinearizationTable",
	Nikon2ColorBalance:              "ColorBalance",
	Nikon2LensData:                  "LensData",
	Nikon2RawImageCentre:            "RawImageCentre",
	Nikon2SensorPixelSize:           "SensorPixelSize",
	Nikon2SceneAssist:               "SceneAssist",
	Nikon2RetouchHistory:            "RetouchHistory",
	Nikon2SerialNumber2:             "SerialNumber2",
	Nikon2ImageDataSize:             "ImageDataSize",
	Nikon2ImageCount:                "ImageCount",
	Nikon2DeletedImageCount:         "DeletedImageCount",
	Nikon2ShutterCount:              "ShutterCount",
	Nikon2FlashInfo:                 "FlashInfo",
	Nikon2ImageOptimization:         "ImageOptimization",
	Nikon2Saturation2:               "Saturation2",
	Nikon2VariProgram:               "VariProgram",
	Nikon2ImageStabilization:        "ImageStabilization",
	Nikon2AFResponse:                "AFResponse",
	Nikon2MultiExposure:             "MultiExposure",
	Nikon2HighISONoiseReduction:     "HighISONoiseReduction",
	Nikon2ToningEffect:              "ToningEffect",
	Nikon2PowerUpTime:               "PowerUpTime",
	Nikon2AFInfo2:                   "AFInfo2",
	Nikon2FileInfo:                  "FileInfo",
	Nikon2AFTune:                    "AFTune",
	Nikon2RetouchInfo:               "RetouchInfo",
	Nikon2PictureControlData2:       "PictureControlData2",
	Nikon2BarometerInfo:             "BarometerInfo",
	Nikon2PrintIM:                   "PrintIM",
	Nikon2NikonCaptureData:          "NikonCaptureData",
	Nikon2NikonCaptureVersion:       "NikonCaptureVersion",
	Nikon2NikonCaptureOffsets:       "NikonCaptureOffsets",
	Nikon2NikonScanIFD:              "NikonScanIFD",
	Nikon2NikonCaptureEditVersions:  "NikonCaptureEditVersions",
	Nikon2NikonICCProfile:           "NikonICCProfile",
	Nikon2NikonCaptureOutput:        "NikonCaptureOutput",
	Nikon2NEFBitDepth:               "NEFBitDepth",
}

// Tags in the "PreviewIFD" sub-IFD that may be found in Nikon2 maker
// notes.  Exiftool: lib/Image/ExifTool/Nikon.pm
// Image::ExifTool::Nikon::PreviewIFD
const (
	Nikon2SubfileType        = 0xFE
	Nikon2Compression        = 0x103
	Nikon2XResolution        = 0x11A
	Nikon2YResolution        = 0x11B
	Nikon2YResolutionUnit    = 0x128
	Nikon2PreviewImageStart  = 0x201
	Nikon2PreviewImageLength = 0x202
	Nikon2YCbCrPositioning   = 0x213
)

// Mapping from Nikon2 Preview IFD tags to strings.
var Nikon2PreviewIFDTagNames = map[tiff.Tag]string{
	Nikon2SubfileType:        "SubfileType",
	Nikon2Compression:        "Compression",
	Nikon2XResolution:        "XResolution",
	Nikon2YResolution:        "YResolution",
	Nikon2YResolutionUnit:    "YResolutionUnit",
	Nikon2PreviewImageStart:  "PreviewImageStart",
	Nikon2PreviewImageLength: "PreviewImageLength",
	Nikon2YCbCrPositioning:   "YCbCrPositioning",
}

// Tags in the "ScanIFD" sub-IFD that may be found in Nikon2 maker
// notes.  Exiftool: lib/Image/ExifTool/Nikon.pm
// Image::ExifTool::Nikon::Scan
const (
	Nikon2FilmType               = 0x2
	Nikon2MultiSample            = 0x40
	Nikon2BitDepth               = 0x41
	Nikon2MasterGain             = 0x50
	Nikon2ColorGain              = 0x51
	Nikon2ScanImageEnhancer      = 0x60
	Nikon2DigitalICE             = 0x100
	Nikon2ROCInfo                = 0x110
	Nikon2GEMInfo                = 0x120
	Nikon2DigitalDEEShadowAdj    = 0x200
	Nikon2DigitalDEEThreshold    = 0x201
	Nikon2DigitalDEEHighlightAdj = 0x202
)

// Mapping from Nikon2 Scan IFD tags to strings.
var Nikon2ScanIFDTagNames = map[tiff.Tag]string{
	Nikon2FilmType:               "FilmType",
	Nikon2MultiSample:            "MultiSample",
	Nikon2BitDepth:               "BitDepth",
	Nikon2MasterGain:             "MasterGain",
	Nikon2ColorGain:              "ColorGain",
	Nikon2ScanImageEnhancer:      "ScanImageEnhancer",
	Nikon2DigitalICE:             "DigitalICE",
	Nikon2ROCInfo:                "ROCInfo",
	Nikon2GEMInfo:                "GEMInfo",
	Nikon2DigitalDEEShadowAdj:    "DigitalDEEShadowAdj",
	Nikon2DigitalDEEThreshold:    "DigitalDEEThreshold",
	Nikon2DigitalDEEHighlightAdj: "DigitalDEEHighlightAdj",
}

// Tags in the Panasonic1 maker note.
// ExifTool lib/Image/ExifTool/Panasonic.pm
const (
	Panasonic1ImageQuality               = 0x1
	Panasonic1FirmwareVersion            = 0x2
	Panasonic1WhiteBalance               = 0x3
	Panasonic1FocusMode                  = 0x7
	Panasonic1AFAreaMode                 = 0xF
	Panasonic1ImageStabilization         = 0x1A
	Panasonic1MacroMode                  = 0x1C
	Panasonic1ShootingMode               = 0x1F
	Panasonic1Audio                      = 0x20
	Panasonic1DataDump                   = 0x21
	Panasonic1WhiteBalanceBias           = 0x23
	Panasonic1FlashBias                  = 0x24
	Panasonic1InternalSerialNumber       = 0x25
	Panasonic1PanasonicExifVersion       = 0x26
	Panasonic1ColorEffect                = 0x28
	Panasonic1TimeSincePowerOn           = 0x29
	Panasonic1BurstMode                  = 0x2A
	Panasonic1SequenceNumber             = 0x2B
	Panasonic1ContrastMode               = 0x2C
	Panasonic1NoiseReduction             = 0x2D
	Panasonic1SelfTimer                  = 0x2E
	Panasonic1Rotation                   = 0x30
	Panasonic1AFAssistLamp               = 0x31
	Panasonic1ColorMode                  = 0x32
	Panasonic1BabyAge1                   = 0x33
	Panasonic1OpticalZoomMode            = 0x34
	Panasonic1ConversionLens             = 0x35
	Panasonic1TravelDay                  = 0x36
	Panasonic1Contrast                   = 0x39
	Panasonic1WorldTimeLocation          = 0x3A
	Panasonic1TextStamp1                 = 0x3B
	Panasonic1ProgramISO                 = 0x3C
	Panasonic1AdvancedSceneType          = 0x3D
	Panasonic1TextStamp2                 = 0x3E
	Panasonic1FacesDetected              = 0x3F
	Panasonic1Saturation                 = 0x40
	Panasonic1Sharpness                  = 0x41
	Panasonic1FilmMode                   = 0x42
	Panasonic1ColorTempKelvin            = 0x44
	Panasonic1BracketSettings            = 0x45
	Panasonic1WBShiftAB                  = 0x46
	Panasonic1WBShiftGM                  = 0x47
	Panasonic1FlashCurtain               = 0x48
	Panasonic1LongExposureNoiseReduction = 0x49
	Panasonic1PanasonicImageWidth        = 0x4B
	Panasonic1PanasonicImageHeight       = 0x4C
	Panasonic1AFPointPosition            = 0x4D
	Panasonic1FaceDetInfo                = 0x4E
	Panasonic1LensType                   = 0x51
	Panasonic1LensSerialNumber           = 0x52
	Panasonic1AccessoryType              = 0x53
	Panasonic1AccessorySerialNumber      = 0x54
	Panasonic1Transform1                 = 0x59
	Panasonic1IntelligentExposure        = 0x5D
	Panasonic1LensFirmwareVersion        = 0x60
	Panasonic1FaceRecInfo                = 0x61
	Panasonic1FlashWarning               = 0x62
	Panasonic1RecognizedFaceFlags        = 0x63
	Panasonic1Title                      = 0x65
	Panasonic1BabyName                   = 0x66
	Panasonic1Location                   = 0x67
	Panasonic1Country                    = 0x69
	Panasonic1State                      = 0x6B
	Panasonic1City1                      = 0x6D
	Panasonic1Landmark                   = 0x6F
	Panasonic1IntelligentResolution      = 0x70
	Panasonic1BurstSpeed                 = 0x77
	Panasonic1IntelligentDRange          = 0x79
	Panasonic1ClearRetouch               = 0x7C
	Panasonic1City2                      = 0x80
	Panasonic1ManometerPressure          = 0x86
	Panasonic1PhotoStyle                 = 0x89
	Panasonic1ShadingCompensation        = 0x8A
	Panasonic1AccelerometerZ             = 0x8C
	Panasonic1AccelerometerX             = 0x8D
	Panasonic1AccelerometerY             = 0x8E
	Panasonic1CameraOrientation          = 0x8F
	Panasonic1RollAngle                  = 0x90
	Panasonic1PitchAngle                 = 0x91
	Panasonic1SweepPanoramaDirection     = 0x93
	Panasonic1SweepPanoramaFieldOfView   = 0x94
	Panasonic1TimerRecording             = 0x96
	Panasonic1InternalNDFilter           = 0x9D
	Panasonic1HDR                        = 0x9E
	Panasonic1ShutterType                = 0x9F
	Panasonic1ClearRetouchValue          = 0xA3
	Panasonic1TouchAE                    = 0xAB
	Panasonic1TimeStamp                  = 0xAF
	Panasonic1PrintIM                    = 0xE00
	Panasonic1MakerNoteVersion           = 0x8000
	Panasonic1SceneMode                  = 0x8001
	Panasonic1WBRedLevel                 = 0x8004
	Panasonic1WBGreenLevel               = 0x8005
	Panasonic1WBBlueLevel                = 0x8006
	Panasonic1FlashFired                 = 0x8007
	Panasonic1TextStamp3                 = 0x8008
	Panasonic1TextStamp4                 = 0x8009
	Panasonic1BabyAge2                   = 0x8010
	Panasonic1Transform2                 = 0x8012
)

// Mapping from Panasonic1 tags to strings.
var Panasonic1TagNames = map[tiff.Tag]string{
	Panasonic1ImageQuality:               "ImageQuality",
	Panasonic1FirmwareVersion:            "FirmwareVersion",
	Panasonic1WhiteBalance:               "WhiteBalance",
	Panasonic1FocusMode:                  "FocusMode",
	Panasonic1AFAreaMode:                 "AFAreaMode",
	Panasonic1ImageStabilization:         "ImageStabilization",
	Panasonic1MacroMode:                  "MacroMode",
	Panasonic1ShootingMode:               "ShootingMode",
	Panasonic1Audio:                      "Audio",
	Panasonic1DataDump:                   "DataDump",
	Panasonic1WhiteBalanceBias:           "WhiteBalanceBias",
	Panasonic1FlashBias:                  "FlashBias",
	Panasonic1InternalSerialNumber:       "InternalSerialNumber",
	Panasonic1PanasonicExifVersion:       "PanasonicExifVersion",
	Panasonic1ColorEffect:                "ColorEffect",
	Panasonic1TimeSincePowerOn:           "TimeSincePowerOn",
	Panasonic1BurstMode:                  "BurstMode",
	Panasonic1SequenceNumber:             "SequenceNumber",
	Panasonic1ContrastMode:               "ContrastMode",
	Panasonic1NoiseReduction:             "NoiseReduction",
	Panasonic1SelfTimer:                  "SelfTimer",
	Panasonic1Rotation:                   "Rotation",
	Panasonic1AFAssistLamp:               "AFAssistLamp",
	Panasonic1ColorMode:                  "ColorMode",
	Panasonic1BabyAge1:                   "BabyAge1",
	Panasonic1OpticalZoomMode:            "OpticalZoomMode",
	Panasonic1ConversionLens:             "ConversionLens",
	Panasonic1TravelDay:                  "TravelDay",
	Panasonic1Contrast:                   "Contrast",
	Panasonic1WorldTimeLocation:          "WorldTimeLocation",
	Panasonic1TextStamp1:                 "TextStamp1",
	Panasonic1ProgramISO:                 "ProgramISO",
	Panasonic1AdvancedSceneType:          "AdvancedSceneType",
	Panasonic1TextStamp2:                 "TextStamp2",
	Panasonic1FacesDetected:              "FacesDetected",
	Panasonic1Saturation:                 "Saturation",
	Panasonic1Sharpness:                  "Sharpness",
	Panasonic1FilmMode:                   "FilmMode",
	Panasonic1ColorTempKelvin:            "ColorTempKelvin",
	Panasonic1BracketSettings:            "BracketSettings",
	Panasonic1WBShiftAB:                  "WBShiftAB",
	Panasonic1WBShiftGM:                  "WBShiftGM",
	Panasonic1FlashCurtain:               "FlashCurtain",
	Panasonic1LongExposureNoiseReduction: "LongExposureNoiseReduction",
	Panasonic1PanasonicImageWidth:        "PanasonicImageWidth",
	Panasonic1PanasonicImageHeight:       "PanasonicImageHeight",
	Panasonic1AFPointPosition:            "AFPointPosition",
	Panasonic1FaceDetInfo:                "FaceDetInfo",
	Panasonic1LensType:                   "LensType",
	Panasonic1LensSerialNumber:           "LensSerialNumber",
	Panasonic1AccessoryType:              "AccessoryType",
	Panasonic1AccessorySerialNumber:      "AccessorySerialNumber",
	Panasonic1Transform1:                 "Transform1",
	Panasonic1IntelligentExposure:        "IntelligentExposure",
	Panasonic1LensFirmwareVersion:        "LensFirmwareVersion",
	Panasonic1FaceRecInfo:                "FaceRecInfo",
	Panasonic1FlashWarning:               "FlashWarning",
	Panasonic1RecognizedFaceFlags:        "RecognizedFaceFlags",
	Panasonic1Title:                      "Title",
	Panasonic1BabyName:                   "BabyName",
	Panasonic1Location:                   "Location",
	Panasonic1Country:                    "Country",
	Panasonic1State:                      "State",
	Panasonic1City1:                      "City1",
	Panasonic1Landmark:                   "Landmark",
	Panasonic1IntelligentResolution:      "IntelligentResolution",
	Panasonic1BurstSpeed:                 "BurstSpeed",
	Panasonic1IntelligentDRange:          "IntelligentD-Range",
	Panasonic1ClearRetouch:               "ClearRetouch",
	Panasonic1City2:                      "City2",
	Panasonic1ManometerPressure:          "ManometerPressure",
	Panasonic1PhotoStyle:                 "PhotoStyle",
	Panasonic1ShadingCompensation:        "ShadingCompensation",
	Panasonic1AccelerometerZ:             "AccelerometerZ",
	Panasonic1AccelerometerX:             "AccelerometerX",
	Panasonic1AccelerometerY:             "AccelerometerY",
	Panasonic1CameraOrientation:          "CameraOrientation",
	Panasonic1RollAngle:                  "RollAngle",
	Panasonic1PitchAngle:                 "PitchAngle",
	Panasonic1SweepPanoramaDirection:     "SweepPanoramaDirection",
	Panasonic1SweepPanoramaFieldOfView:   "SweepPanoramaFieldOfView",
	Panasonic1TimerRecording:             "TimerRecording",
	Panasonic1InternalNDFilter:           "InternalNDFilter",
	Panasonic1HDR:                        "HDR",
	Panasonic1ShutterType:                "ShutterType",
	Panasonic1ClearRetouchValue:          "ClearRetouchValue",
	Panasonic1TouchAE:                    "TouchAE",
	Panasonic1TimeStamp:                  "TimeStamp",
	Panasonic1PrintIM:                    "PrintIM",
	Panasonic1MakerNoteVersion:           "MakerNoteVersion",
	Panasonic1SceneMode:                  "SceneMode",
	Panasonic1WBRedLevel:                 "WBRedLevel",
	Panasonic1WBGreenLevel:               "WBGreenLevel",
	Panasonic1WBBlueLevel:                "WBBlueLevel",
	Panasonic1FlashFired:                 "FlashFired",
	Panasonic1TextStamp3:                 "TextStamp3",
	Panasonic1TextStamp4:                 "TextStamp4",
	Panasonic1BabyAge2:                   "BabyAge2",
	Panasonic1Transform2:                 "Transform2",
}
