package exif44

import (
	tiff "github.com/garyhouston/tiff66"
)

// Tag names are from ExifTool 10.63, which also has information about
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

// Tags in the Fujifilm1 maker note.
const (
	Fujifilm1Version                 = 0x0
	Fujifilm1InternalSerialNumber    = 0x10
	Fujifilm1Quality                 = 0x1000
	Fujifilm1Sharpness               = 0x1001
	Fujifilm1WhiteBalance            = 0x1002
	Fujifilm1Saturation              = 0x1003
	Fujifilm1Contrast1               = 0x1004
	Fujifilm1ColorTemperature        = 0x1005
	Fujifilm1Contrast2               = 0x1006
	Fujifilm1WhiteBalanceFineTune    = 0x100A
	Fujifilm1NoiseReduction          = 0x100B
	Fujifilm1HighISONoiseReduction   = 0x100E
	Fujifilm1FujiFlashMode           = 0x1010
	Fujifilm1FlashExposureComp       = 0x1011
	Fujifilm1Macro                   = 0x1020
	Fujifilm1FocusMode               = 0x1021
	Fujifilm1AFMode                  = 0x1022
	Fujifilm1FocusPixel              = 0x1023
	Fujifilm1SlowSync                = 0x1030
	Fujifilm1PictureMode             = 0x1031
	Fujifilm1ExposureCount           = 0x1032
	Fujifilm1EXRAuto                 = 0x1033
	Fujifilm1EXRMode                 = 0x1034
	Fujifilm1ShadowTone              = 0x1040
	Fujifilm1HighlightTone           = 0x1041
	Fujifilm1DigitalZoom             = 0x1044
	Fujifilm1ShutterType             = 0x1050
	Fujifilm1AutoBracketing          = 0x1100
	Fujifilm1SequenceNumber          = 0x1101
	Fujifilm1PanoramaAngle           = 0x1153
	Fujifilm1PanoramaDirection       = 0x1154
	Fujifilm1AdvancedFilter          = 0x1201
	Fujifilm1ColorMode               = 0x1210
	Fujifilm1BlurWarning             = 0x1300
	Fujifilm1FocusWarning            = 0x1301
	Fujifilm1ExposureWarning         = 0x1302
	Fujifilm1GEImageSize             = 0x1304
	Fujifilm1DynamicRange            = 0x1400
	Fujifilm1FilmMode                = 0x1401
	Fujifilm1DynamicRangeSetting     = 0x1402
	Fujifilm1DevelopmentDynamicRange = 0x1403
	Fujifilm1MinFocalLength          = 0x1404
	Fujifilm1MaxFocalLength          = 0x1405
	Fujifilm1MaxApertureAtMinFocal   = 0x1406
	Fujifilm1MaxApertureAtMaxFocal   = 0x1407
	Fujifilm1AutoDynamicRange        = 0x140B
	Fujifilm1ImageStabilization      = 0x1422
	Fujifilm1Rating                  = 0x1431
	Fujifilm1ImageGeneration         = 0x1436
	Fujifilm1ImageCount              = 0x1438
	Fujifilm1FrameRate               = 0x3820
	Fujifilm1FrameWidth              = 0x3821
	Fujifilm1FrameHeight             = 0x3822
	Fujifilm1FacesDetected           = 0x4100
	Fujifilm1FacePositions           = 0x4103
	Fujifilm1FaceRecInfo             = 0x4282
	Fujifilm1FileSource              = 0x8000
	Fujifilm1OrderNumber             = 0x8002
	Fujifilm1FrameNumber             = 0x8003
	Fujifilm1Parallax                = 0xB211
)

// Mapping from Fujifilm1 tags to strings.
var Fujifilm1TagNames = map[tiff.Tag]string{
	Fujifilm1Version:                 "Version",
	Fujifilm1InternalSerialNumber:    "InternalSerialNumber",
	Fujifilm1Quality:                 "Quality",
	Fujifilm1Sharpness:               "Sharpness",
	Fujifilm1WhiteBalance:            "WhiteBalance",
	Fujifilm1Saturation:              "Saturation",
	Fujifilm1Contrast1:               "Contrast1",
	Fujifilm1ColorTemperature:        "ColorTemperature",
	Fujifilm1Contrast2:               "Contrast2",
	Fujifilm1WhiteBalanceFineTune:    "WhiteBalanceFineTune",
	Fujifilm1NoiseReduction:          "NoiseReduction",
	Fujifilm1HighISONoiseReduction:   "HighISONoiseReduction",
	Fujifilm1FujiFlashMode:           "FujiFlashMode",
	Fujifilm1FlashExposureComp:       "FlashExposureComp",
	Fujifilm1Macro:                   "Macro",
	Fujifilm1FocusMode:               "FocusMode",
	Fujifilm1AFMode:                  "AFMode",
	Fujifilm1FocusPixel:              "FocusPixel",
	Fujifilm1SlowSync:                "SlowSync",
	Fujifilm1PictureMode:             "PictureMode",
	Fujifilm1ExposureCount:           "ExposureCount",
	Fujifilm1EXRAuto:                 "EXRAuto",
	Fujifilm1EXRMode:                 "EXRMode",
	Fujifilm1ShadowTone:              "ShadowTone",
	Fujifilm1HighlightTone:           "HighlightTone",
	Fujifilm1DigitalZoom:             "DigitalZoom",
	Fujifilm1ShutterType:             "ShutterType",
	Fujifilm1AutoBracketing:          "AutoBracketing",
	Fujifilm1SequenceNumber:          "SequenceNumber",
	Fujifilm1PanoramaAngle:           "PanoramaAngle",
	Fujifilm1PanoramaDirection:       "PanoramaDirection",
	Fujifilm1AdvancedFilter:          "AdvancedFilter",
	Fujifilm1ColorMode:               "ColorMode",
	Fujifilm1BlurWarning:             "BlurWarning",
	Fujifilm1FocusWarning:            "FocusWarning",
	Fujifilm1ExposureWarning:         "ExposureWarning",
	Fujifilm1GEImageSize:             "GEImageSize",
	Fujifilm1DynamicRange:            "DynamicRange",
	Fujifilm1FilmMode:                "FilmMode",
	Fujifilm1DynamicRangeSetting:     "DynamicRangeSetting",
	Fujifilm1DevelopmentDynamicRange: "DevelopmentDynamicRange",
	Fujifilm1MinFocalLength:          "MinFocalLength",
	Fujifilm1MaxFocalLength:          "MaxFocalLength",
	Fujifilm1MaxApertureAtMinFocal:   "MaxApertureAtMinFocal",
	Fujifilm1MaxApertureAtMaxFocal:   "MaxApertureAtMaxFocal",
	Fujifilm1AutoDynamicRange:        "AutoDynamicRange",
	Fujifilm1ImageStabilization:      "ImageStabilization",
	Fujifilm1Rating:                  "Rating",
	Fujifilm1ImageGeneration:         "ImageGeneration",
	Fujifilm1ImageCount:              "ImageCount",
	Fujifilm1FrameRate:               "FrameRate",
	Fujifilm1FrameWidth:              "FrameWidth",
	Fujifilm1FrameHeight:             "FrameHeight",
	Fujifilm1FacesDetected:           "FacesDetected",
	Fujifilm1FacePositions:           "FacePositions",
	Fujifilm1FaceRecInfo:             "FaceRecInfo",
	Fujifilm1FileSource:              "FileSource",
	Fujifilm1OrderNumber:             "OrderNumber",
	Fujifilm1FrameNumber:             "FrameNumber",
	Fujifilm1Parallax:                "Parallax",
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

// Tags in the Olympus1 maker note.
// ExifTool lib/Image/ExifTool/Olympus.pm ::Main
const (
	Olympus1MakerNoteVersion         = 0x0
	Olympus1MinoltaCameraSettingsOld = 0x1
	Olympus1MinoltaCameraSettings    = 0x3
	Olympus1CompressedImageSize      = 0x40
	Olympus1PreviewImageData         = 0x81
	Olympus1PreviewImageStart        = 0x88
	Olympus1PreviewImageLength       = 0x89
	Olympus1ThumbnailImage           = 0x100
	Olympus1BodyFirmwareVersion      = 0x104
	Olympus1SpecialMode              = 0x200
	Olympus1Quality                  = 0x201
	Olympus1Macro                    = 0x202
	Olympus1BWMode                   = 0x203
	Olympus1DigitalZoom              = 0x204
	Olympus1FocalPlaneDiagonal       = 0x205
	Olympus1LensDistortionParams     = 0x206
	Olympus1CameraType               = 0x207
	Olympus1TextInfo                 = 0x208
	Olympus1CameraID                 = 0x209
	Olympus1EpsonImageWidth          = 0x20B
	Olympus1EpsonImageHeight         = 0x20C
	Olympus1EpsonSoftware            = 0x20D
	Olympus1Preview                  = 0x280
	Olympus1PreCaptureFrames         = 0x300
	Olympus1WhiteBoard               = 0x301
	Olympus1OneTouchWB               = 0x302
	Olympus1WhiteBalanceBracket      = 0x303
	Olympus1WhiteBalanceBias         = 0x304
	Olympus1BlackLevel               = 0x401
	Olympus1SceneMode                = 0x403
	Olympus1SerialNumber             = 0x404
	Olympus1Firmware                 = 0x405
	Olympus1PrintIM                  = 0xE00
	Olympus1DataDump                 = 0xF00
	Olympus1DataDump2                = 0xF01
	Olympus1ZoomedPreviewStart       = 0xF04
	Olympus1ZoomedPreviewLength      = 0xF05
	Olympus1ZoomedPreviewSize        = 0xF06
	Olympus1ShutterSpeedValue        = 0x1000
	Olympus1ISOValue                 = 0x1001
	Olympus1ApertureValue            = 0x1002
	Olympus1BrightnessValue          = 0x1003
	Olympus1FlashMode                = 0x1004
	Olympus1FlashDevice              = 0x1005
	Olympus1ExposureCompensation     = 0x1006
	Olympus1SensorTemperature        = 0x1007
	Olympus1LensTemperature          = 0x1008
	Olympus1LightCondition           = 0x1009
	Olympus1FocusRange               = 0x100A
	Olympus1FocusMode                = 0x100B
	Olympus1ManualFocusDistance      = 0x100C
	Olympus1ZoomStepCount            = 0x100D
	Olympus1FocusStepCount           = 0x100E
	Olympus1Sharpness                = 0x100F
	Olympus1FlashChargeLevel         = 0x1010
	Olympus1ColorMatrix              = 0x1011
	Olympus1BlackLevel2              = 0x1012
	Olympus1ColorTemperatureBG       = 0x1013
	Olympus1ColorTemperatureRG       = 0x1014
	Olympus1WBMode                   = 0x1015
	Olympus1RedBalance               = 0x1017
	Olympus1BlueBalance              = 0x1018
	Olympus1ColorMatrixNumber        = 0x1019
	Olympus1SerialNumber2            = 0x101A
	Olympus1ExternalFlashAE1_0       = 0x101B
	Olympus1ExternalFlashAE2_0       = 0x101C
	Olympus1InternalFlashAE1_0       = 0x101D
	Olympus1InternalFlashAE2_0       = 0x101E
	Olympus1ExternalFlashAE1         = 0x101F
	Olympus1ExternalFlashAE2         = 0x1020
	Olympus1InternalFlashAE1         = 0x1021
	Olympus1InternalFlashAE2         = 0x1022
	Olympus1FlashExposureComp        = 0x1023
	Olympus1InternalFlashTable       = 0x1024
	Olympus1ExternalFlashGValue      = 0x1025
	Olympus1ExternalFlashBounce      = 0x1026
	Olympus1ExternalFlashZoom        = 0x1027
	Olympus1ExternalFlashMode        = 0x1028
	Olympus1Contrast                 = 0x1029
	Olympus1SharpnessFactor          = 0x102A
	Olympus1ColorControl             = 0x102B
	Olympus1ValidBits                = 0x102C
	Olympus1CoringFilter             = 0x102D
	Olympus1OlympusImageWidth        = 0x102E
	Olympus1OlympusImageHeight       = 0x102F
	Olympus1SceneDetect              = 0x1030
	Olympus1SceneArea                = 0x1031
	Olympus1SceneDetectData          = 0x1033
	Olympus1CompressionRatio         = 0x1034
	Olympus1PreviewImageValid        = 0x1035
	Olympus1PreviewImageStart2       = 0x1036
	Olympus1PreviewImageLength2      = 0x1037
	Olympus1AFResult                 = 0x1038
	Olympus1CCDScanMode              = 0x1039
	Olympus1NoiseReduction           = 0x103A
	Olympus1FocusStepInfinity        = 0x103B
	Olympus1FocusStepNear            = 0x103C
	Olympus1LightValueCenter         = 0x103D
	Olympus1LightValuePeriphery      = 0x103E
	Olympus1FieldCount               = 0x103F
	Olympus1EquipmentIFD             = 0x2010
	Olympus1CameraSettingsIFD        = 0x2020
	Olympus1RawDevelopmentIFD        = 0x2030
	Olympus1RawDev2IFD               = 0x2031
	Olympus1ImageProcessingIFD       = 0x2040
	Olympus1FocusInfoIFD             = 0x2050
	Olympus1RawInfoIFD               = 0x3000
)

// Mapping from Olympus1 tags to strings.
var Olympus1TagNames = map[tiff.Tag]string{
	Olympus1MakerNoteVersion:         "MakerNoteVersion",
	Olympus1MinoltaCameraSettingsOld: "MinoltaCameraSettingsOld",
	Olympus1MinoltaCameraSettings:    "MinoltaCameraSettings",
	Olympus1CompressedImageSize:      "CompressedImageSize",
	Olympus1PreviewImageData:         "PreviewImageData",
	Olympus1PreviewImageStart:        "PreviewImageStart",
	Olympus1PreviewImageLength:       "PreviewImageLength",
	Olympus1ThumbnailImage:           "ThumbnailImage",
	Olympus1BodyFirmwareVersion:      "BodyFirmwareVersion",
	Olympus1SpecialMode:              "SpecialMode",
	Olympus1Quality:                  "Quality",
	Olympus1Macro:                    "Macro",
	Olympus1BWMode:                   "BWMode",
	Olympus1DigitalZoom:              "DigitalZoom",
	Olympus1FocalPlaneDiagonal:       "FocalPlaneDiagonal",
	Olympus1LensDistortionParams:     "LensDistortionParams",
	Olympus1CameraType:               "CameraType",
	Olympus1TextInfo:                 "TextInfo",
	Olympus1CameraID:                 "CameraID",
	Olympus1EpsonImageWidth:          "EpsonImageWidth",
	Olympus1EpsonImageHeight:         "EpsonImageHeight",
	Olympus1EpsonSoftware:            "EpsonSoftware",
	Olympus1Preview:                  "Preview",
	Olympus1PreCaptureFrames:         "PreCaptureFrames",
	Olympus1WhiteBoard:               "WhiteBoard",
	Olympus1OneTouchWB:               "OneTouchWB",
	Olympus1WhiteBalanceBracket:      "WhiteBalanceBracket",
	Olympus1WhiteBalanceBias:         "WhiteBalanceBias",
	Olympus1BlackLevel:               "BlackLevel",
	Olympus1SceneMode:                "SceneMode",
	Olympus1SerialNumber:             "SerialNumber",
	Olympus1Firmware:                 "Firmware",
	Olympus1PrintIM:                  "PrintIM",
	Olympus1DataDump:                 "DataDump",
	Olympus1DataDump2:                "DataDump2",
	Olympus1ZoomedPreviewStart:       "ZoomedPreviewStart",
	Olympus1ZoomedPreviewLength:      "ZoomedPreviewLength",
	Olympus1ZoomedPreviewSize:        "ZoomedPreviewSize",
	Olympus1ShutterSpeedValue:        "ShutterSpeedValue",
	Olympus1ISOValue:                 "ISOValue",
	Olympus1ApertureValue:            "ApertureValue",
	Olympus1BrightnessValue:          "BrightnessValue",
	Olympus1FlashMode:                "FlashMode",
	Olympus1FlashDevice:              "FlashDevice",
	Olympus1ExposureCompensation:     "ExposureCompensation",
	Olympus1SensorTemperature:        "SensorTemperature",
	Olympus1LensTemperature:          "LensTemperature",
	Olympus1LightCondition:           "LightCondition",
	Olympus1FocusRange:               "FocusRange",
	Olympus1FocusMode:                "FocusMode",
	Olympus1ManualFocusDistance:      "ManualFocusDistance",
	Olympus1ZoomStepCount:            "ZoomStepCount",
	Olympus1FocusStepCount:           "FocusStepCount",
	Olympus1Sharpness:                "Sharpness",
	Olympus1FlashChargeLevel:         "FlashChargeLevel",
	Olympus1ColorMatrix:              "ColorMatrix",
	Olympus1BlackLevel2:              "BlackLevel2",
	Olympus1ColorTemperatureBG:       "ColorTemperatureBG",
	Olympus1ColorTemperatureRG:       "ColorTemperatureRG",
	Olympus1WBMode:                   "WBMode",
	Olympus1RedBalance:               "RedBalance",
	Olympus1BlueBalance:              "BlueBalance",
	Olympus1ColorMatrixNumber:        "ColorMatrixNumber",
	Olympus1SerialNumber2:            "SerialNumber2",
	Olympus1ExternalFlashAE1_0:       "ExternalFlashAE1_0",
	Olympus1ExternalFlashAE2_0:       "ExternalFlashAE2_0",
	Olympus1InternalFlashAE1_0:       "InternalFlashAE1_0",
	Olympus1InternalFlashAE2_0:       "InternalFlashAE2_0",
	Olympus1ExternalFlashAE1:         "ExternalFlashAE1",
	Olympus1ExternalFlashAE2:         "ExternalFlashAE2",
	Olympus1InternalFlashAE1:         "InternalFlashAE1",
	Olympus1InternalFlashAE2:         "InternalFlashAE2",
	Olympus1FlashExposureComp:        "FlashExposureComp",
	Olympus1InternalFlashTable:       "InternalFlashTable",
	Olympus1ExternalFlashGValue:      "ExternalFlashGValue",
	Olympus1ExternalFlashBounce:      "ExternalFlashBounce",
	Olympus1ExternalFlashZoom:        "ExternalFlashZoom",
	Olympus1ExternalFlashMode:        "ExternalFlashMode",
	Olympus1Contrast:                 "Contrast",
	Olympus1SharpnessFactor:          "SharpnessFactor",
	Olympus1ColorControl:             "ColorControl",
	Olympus1ValidBits:                "ValidBits",
	Olympus1CoringFilter:             "CoringFilter",
	Olympus1OlympusImageWidth:        "OlympusImageWidth",
	Olympus1OlympusImageHeight:       "OlympusImageHeight",
	Olympus1SceneDetect:              "SceneDetect",
	Olympus1SceneArea:                "SceneArea",
	Olympus1SceneDetectData:          "SceneDetectData",
	Olympus1CompressionRatio:         "CompressionRatio",
	Olympus1PreviewImageValid:        "PreviewImageValid",
	Olympus1PreviewImageStart2:       "PreviewImageStart2",
	Olympus1PreviewImageLength2:      "PreviewImageLength2",
	Olympus1AFResult:                 "AFResult",
	Olympus1CCDScanMode:              "CCDScanMode",
	Olympus1NoiseReduction:           "NoiseReduction",
	Olympus1FocusStepInfinity:        "FocusStepInfinity",
	Olympus1FocusStepNear:            "FocusStepNear",
	Olympus1LightValueCenter:         "LightValueCenter",
	Olympus1LightValuePeriphery:      "LightValuePeriphery",
	Olympus1FieldCount:               "FieldCount",
	Olympus1EquipmentIFD:             "EquipmentIFD",
	Olympus1CameraSettingsIFD:        "CameraSettingsIFD",
	Olympus1RawDevelopmentIFD:        "RawDevelopmentIFD",
	Olympus1RawDev2IFD:               "RawDev2IFD",
	Olympus1ImageProcessingIFD:       "ImageProcessingIFD",
	Olympus1FocusInfoIFD:             "FocusInfoIFD",
	Olympus1RawInfoIFD:               "RawInfoIFD",
}

// Tags in the Olympus1 Equipment IFD
// ExifTool lib/Image/ExifTool/Olympus.pm ::Equipment
const (
	Olympus1EqVersion                 = 0x0
	Olympus1EqCameraType2             = 0x100
	Olympus1EqSerialNumber            = 0x101
	Olympus1EqInternalSerialNumber    = 0x102
	Olympus1EqFocalPlaneDiagonal      = 0x103
	Olympus1EqBodyFirmwareVersion     = 0x104
	Olympus1EqLensType                = 0x201
	Olympus1EqLensSerialNumber        = 0x202
	Olympus1EqLensModel               = 0x203
	Olympus1EqLensFirmwareVersion     = 0x204
	Olympus1EqMaxApertureAtMinFocal   = 0x205
	Olympus1EqMaxApertureAtMaxFocal   = 0x206
	Olympus1EqMinFocalLength          = 0x207
	Olympus1EqMaxFocalLength          = 0x208
	Olympus1EqMaxAperture             = 0x20A
	Olympus1EqLensProperties          = 0x20B
	Olympus1EqExtender                = 0x300
	Olympus1EqExtenderSerialNumber    = 0x302
	Olympus1EqExtenderModel           = 0x303
	Olympus1EqExtenderFirmwareVersion = 0x304
	Olympus1EqConversionLens          = 0x403
	Olympus1EqFlashType               = 0x1000
	Olympus1EqFlashModel              = 0x1001
	Olympus1EqFlashFirmwareVersion    = 0x1002
	Olympus1EqFlashSerialNumber       = 0x1003
)

// Mapping from Olympus1 Equipment tags to strings.
var Olympus1EquipmentTagNames = map[tiff.Tag]string{
	Olympus1EqVersion:                 "Version",
	Olympus1EqCameraType2:             "CameraType2",
	Olympus1EqSerialNumber:            "SerialNumber",
	Olympus1EqInternalSerialNumber:    "InternalSerialNumber",
	Olympus1EqFocalPlaneDiagonal:      "FocalPlaneDiagonal",
	Olympus1EqBodyFirmwareVersion:     "BodyFirmwareVersion",
	Olympus1EqLensType:                "LensType",
	Olympus1EqLensSerialNumber:        "LensSerialNumber",
	Olympus1EqLensModel:               "LensModel",
	Olympus1EqLensFirmwareVersion:     "LensFirmwareVersion",
	Olympus1EqMaxApertureAtMinFocal:   "MaxApertureAtMinFocal",
	Olympus1EqMaxApertureAtMaxFocal:   "MaxApertureAtMaxFocal",
	Olympus1EqMinFocalLength:          "MinFocalLength",
	Olympus1EqMaxFocalLength:          "MaxFocalLength",
	Olympus1EqMaxAperture:             "MaxAperture",
	Olympus1EqLensProperties:          "LensProperties",
	Olympus1EqExtender:                "Extender",
	Olympus1EqExtenderSerialNumber:    "ExtenderSerialNumber",
	Olympus1EqExtenderModel:           "ExtenderModel",
	Olympus1EqExtenderFirmwareVersion: "ExtenderFirmwareVersion",
	Olympus1EqConversionLens:          "ConversionLens",
	Olympus1EqFlashType:               "FlashType",
	Olympus1EqFlashModel:              "FlashModel",
	Olympus1EqFlashFirmwareVersion:    "FlashFirmwareVersion",
	Olympus1EqFlashSerialNumber:       "FlashSerialNumber",
}

// Tags in the Olympus1 Camera Settings IFD
// ExifTool lib/Image/ExifTool/Olympus.pm ::CameraSettings
const (
	Olympus1CSVersion                   = 0x0
	Olympus1CSPreviewImageValid         = 0x100
	Olympus1CSPreviewImageStart         = 0x101
	Olympus1CSPreviewImageLength        = 0x102
	Olympus1CSExposureMode              = 0x200
	Olympus1CSAELock                    = 0x201
	Olympus1CSMeteringMode              = 0x202
	Olympus1CSExposureShift             = 0x203
	Olympus1CSNDFilter                  = 0x204
	Olympus1CSMacroMode                 = 0x300
	Olympus1CSFocusMode                 = 0x301
	Olympus1CSFocusProcess              = 0x302
	Olympus1CSAFSearch                  = 0x303
	Olympus1CSAFAreas                   = 0x304
	Olympus1CSAFPointSelected           = 0x305
	Olympus1CSAFFineTune                = 0x306
	Olympus1CSAFFineTuneAdj             = 0x307
	Olympus1CSFlashMode                 = 0x400
	Olympus1CSFlashExposureComp         = 0x401
	Olympus1CSFlashRemoteControl        = 0x403
	Olympus1CSFlashControlMode          = 0x404
	Olympus1CSFlashIntensity            = 0x405
	Olympus1CSManualFlashStrength       = 0x406
	Olympus1CSWhiteBalance2             = 0x500
	Olympus1CSWhiteBalanceTemperature   = 0x501
	Olympus1CSWhiteBalanceBracket       = 0x502
	Olympus1CSCustomSaturation          = 0x503
	Olympus1CSModifiedSaturation        = 0x504
	Olympus1CSConstrastSetting          = 0x505
	Olympus1CSSharpnessSetting          = 0x506
	Olympus1CSColorSpace                = 0x507
	Olympus1CSSceneMode                 = 0x509
	Olympus1CSNoiseReduction            = 0x50A
	Olympus1CSDistortionCorrection      = 0x50B
	Olympus1CSShadingCompensation       = 0x50C
	Olympus1CSCompressionFactor         = 0x50D
	Olympus1CSGradation                 = 0x50F
	Olympus1CSPictureMode               = 0x520
	Olympus1CSPictureModeSaturation     = 0x521
	Olympus1CSPictureModeHue            = 0x522
	Olympus1CSPictureModeContrast       = 0x523
	Olympus1CSPictureModeSharpness      = 0x524
	Olympus1CSPictureModeBWFilter       = 0x525
	Olympus1CSPictureModeTone           = 0x526
	Olympus1CSNoiseFilter               = 0x527
	Olympus1CSArtFilter                 = 0x529
	Olympus1CSMagicFilter               = 0x52C
	Olympus1CSPictureModeEffect         = 0x52D
	Olympus1CSToneLevel                 = 0x52E
	Olympus1CSArtFilterEffect           = 0x52F
	Olympus1CSColorCreatorEffect        = 0x532
	Olympus1CSMonochromeProfileSettings = 0x537
	Olympus1CSFilmGrainEffect           = 0x538
	Olympus1CSColorProfileSettings      = 0x539
	Olympus1CSMonochromeVignetting      = 0x53A
	Olympus1CSMonochromeColor           = 0x53B
	Olympus1CSDriveMode                 = 0x600
	Olympus1CSPanoramaMode              = 0x601
	Olympus1CSImageQuality2             = 0x603
	Olympus1CSImageStablization         = 0x604
	Olympus1CSStackedImage              = 0x804
	Olympus1CSManometerPressure         = 0x900
	Olympus1CSManometerReading          = 0x901
	Olympus1CSExtendedWBDetect          = 0x902
	Olympus1CSRollAngle                 = 0x903
	Olympus1CSPitchAngle                = 0x904
	Olympus1CSDateTimeUTC               = 0x908
)

// Mapping from Olympus1 Camera Settings tags to strings.
var Olympus1CameraSettingsTagNames = map[tiff.Tag]string{
	Olympus1CSVersion:                   "Version",
	Olympus1CSPreviewImageValid:         "PreviewImageValid",
	Olympus1CSPreviewImageStart:         "PreviewImageStart",
	Olympus1CSPreviewImageLength:        "PreviewImageLength",
	Olympus1CSExposureMode:              "ExposureMode",
	Olympus1CSAELock:                    "AELock",
	Olympus1CSMeteringMode:              "MeteringMode",
	Olympus1CSExposureShift:             "ExposureShift",
	Olympus1CSNDFilter:                  "NDFilter",
	Olympus1CSMacroMode:                 "MacroMode",
	Olympus1CSFocusMode:                 "FocusMode",
	Olympus1CSFocusProcess:              "FocusProcess",
	Olympus1CSAFSearch:                  "AFSearch",
	Olympus1CSAFAreas:                   "AFAreas",
	Olympus1CSAFPointSelected:           "AFPointSelected",
	Olympus1CSAFFineTune:                "AFFineTune",
	Olympus1CSAFFineTuneAdj:             "AFFineTuneAdj",
	Olympus1CSFlashMode:                 "FlashMode",
	Olympus1CSFlashExposureComp:         "FlashExposureComp",
	Olympus1CSFlashRemoteControl:        "FlashRemoteControl",
	Olympus1CSFlashControlMode:          "FlashControlMode",
	Olympus1CSFlashIntensity:            "FlashIntensity",
	Olympus1CSManualFlashStrength:       "ManualFlashStrength",
	Olympus1CSWhiteBalance2:             "WhiteBalance2",
	Olympus1CSWhiteBalanceTemperature:   "WhiteBalanceTemperature",
	Olympus1CSWhiteBalanceBracket:       "WhiteBalanceBracket",
	Olympus1CSCustomSaturation:          "CustomSaturation",
	Olympus1CSModifiedSaturation:        "ModifiedSaturation",
	Olympus1CSConstrastSetting:          "ConstrastSetting",
	Olympus1CSSharpnessSetting:          "SharpnessSetting",
	Olympus1CSColorSpace:                "ColorSpace",
	Olympus1CSSceneMode:                 "SceneMode",
	Olympus1CSNoiseReduction:            "NoiseReduction",
	Olympus1CSDistortionCorrection:      "DistortionCorrection",
	Olympus1CSShadingCompensation:       "ShadingCompensation",
	Olympus1CSCompressionFactor:         "CompressionFactor",
	Olympus1CSGradation:                 "Gradation",
	Olympus1CSPictureMode:               "PictureMode",
	Olympus1CSPictureModeSaturation:     "PictureModeSaturation",
	Olympus1CSPictureModeHue:            "PictureModeHue",
	Olympus1CSPictureModeContrast:       "PictureModeContrast",
	Olympus1CSPictureModeSharpness:      "PictureModeSharpness",
	Olympus1CSPictureModeBWFilter:       "PictureModeBWFilter",
	Olympus1CSPictureModeTone:           "PictureModeTone",
	Olympus1CSNoiseFilter:               "NoiseFilter",
	Olympus1CSArtFilter:                 "ArtFilter",
	Olympus1CSMagicFilter:               "MagicFilter",
	Olympus1CSPictureModeEffect:         "PictureModeEffect",
	Olympus1CSToneLevel:                 "ToneLevel",
	Olympus1CSArtFilterEffect:           "ArtFilterEffect",
	Olympus1CSColorCreatorEffect:        "ColorCreatorEffect",
	Olympus1CSMonochromeProfileSettings: "MonochromeProfileSettings",
	Olympus1CSFilmGrainEffect:           "FilmGrainEffect",
	Olympus1CSColorProfileSettings:      "ColorProfileSettings",
	Olympus1CSMonochromeVignetting:      "MonochromeVignetting",
	Olympus1CSMonochromeColor:           "MonochromeColor",
	Olympus1CSDriveMode:                 "DriveMode",
	Olympus1CSPanoramaMode:              "PanoramaMode",
	Olympus1CSImageQuality2:             "ImageQuality2",
	Olympus1CSImageStablization:         "ImageStablization",
	Olympus1CSStackedImage:              "StackedImage",
	Olympus1CSManometerPressure:         "ManometerPressure",
	Olympus1CSManometerReading:          "ManometerReading",
	Olympus1CSExtendedWBDetect:          "ExtendedWBDetect",
	Olympus1CSRollAngle:                 "RollAngle",
	Olympus1CSPitchAngle:                "PitchAngle",
	Olympus1CSDateTimeUTC:               "DateTimeUTC",
}

// Tags in the Olympus1 Raw Development IFD
// ExifTool lib/Image/ExifTool/Olympus.pm ::RawDevelopment
const (
	Olympus1RDVersion             = 0x0
	Olympus1RDExposureBiasValue   = 0x100
	Olympus1RDWhiteBalanceValue   = 0x101
	Olympus1RDWBFineAdjustment    = 0x102
	Olympus1RDGrayPoint           = 0x103
	Olympus1RDSaturationEmphasis  = 0x104
	Olympus1RDMemoryColorEmphasis = 0x105
	Olympus1RDContrastValue       = 0x106
	Olympus1RDSharpnessValue      = 0x107
	Olympus1RDColorSpace          = 0x108
	Olympus1RDEngine              = 0x109
	Olympus1RDNoiseReduction      = 0x10A
	Olympus1RDEditStatus          = 0x10B
	Olympus1RDSettings            = 0x10C
)

// Mapping from Olympus1 Raw Development tags to strings.
var Olympus1RawDevelopmentTagNames = map[tiff.Tag]string{
	Olympus1RDVersion:             "Version",
	Olympus1RDExposureBiasValue:   "ExposureBiasValue",
	Olympus1RDWhiteBalanceValue:   "WhiteBalanceValue",
	Olympus1RDWBFineAdjustment:    "WBFineAdjustment",
	Olympus1RDGrayPoint:           "GrayPoint",
	Olympus1RDSaturationEmphasis:  "SaturationEmphasis",
	Olympus1RDMemoryColorEmphasis: "MemoryColorEmphasis",
	Olympus1RDContrastValue:       "ContrastValue",
	Olympus1RDSharpnessValue:      "SharpnessValue",
	Olympus1RDColorSpace:          "ColorSpace",
	Olympus1RDEngine:              "Engine",
	Olympus1RDNoiseReduction:      "NoiseReduction",
	Olympus1RDEditStatus:          "EditStatus",
	Olympus1RDSettings:            "Settings",
}

// Tags in the Olympus1 Raw Dev 2 IFD
// ExifTool lib/Image/ExifTool/Olympus.pm ::RawDevelopment2
const (
	Olympus1RD2Version             = 0x0
	Olympus1RD2ExposureBiasValue   = 0x100
	Olympus1RD2WhiteBalance        = 0x101
	Olympus1RD2WhiteBalanceValue   = 0x102
	Olympus1RD2WBFineAdjustment    = 0x103
	Olympus1RD2GrayPoint           = 0x104
	Olympus1RD2ContrastValue       = 0x105
	Olympus1RD2SharpnessValue      = 0x106
	Olympus1RD2SaturationEmphasis  = 0x107
	Olympus1RD2MemoryColorEmphasis = 0x108
	Olympus1RD2ColorSpace          = 0x109
	Olympus1RD2NoiseReduction      = 0x10A
	Olympus1RD2Engine              = 0x10B
	Olympus1RD2PictureMode         = 0x10C
	Olympus1RD2PMSaturation        = 0x10D
	Olympus1RD2PMContrast          = 0x10E
	Olympus1RD2PMSharpness         = 0x10F
	Olympus1RD2PM_BWFilter         = 0x110
	Olympus1RD2PMPictureTone       = 0x111
	Olympus1RD2DevGradation        = 0x112
	Olympus1RD2DevSaturation3      = 0x113
	Olympus1RD2AutoGradation       = 0x119
	Olympus1RD2PMNoiseFilter       = 0x120
	Olympus1RD2ArtFilter           = 0x121
)

// Mapping from Olympus1 Raw Development tags to strings.
var Olympus1RawDev2TagNames = map[tiff.Tag]string{
	Olympus1RD2Version:             "Version",
	Olympus1RD2ExposureBiasValue:   "ExposureBiasValue",
	Olympus1RD2WhiteBalance:        "WhiteBalance",
	Olympus1RD2WhiteBalanceValue:   "WhiteBalanceValue",
	Olympus1RD2WBFineAdjustment:    "WBFineAdjustment",
	Olympus1RD2GrayPoint:           "GrayPoint",
	Olympus1RD2ContrastValue:       "ContrastValue",
	Olympus1RD2SharpnessValue:      "SharpnessValue",
	Olympus1RD2SaturationEmphasis:  "SaturationEmphasis",
	Olympus1RD2MemoryColorEmphasis: "MemoryColorEmphasis",
	Olympus1RD2ColorSpace:          "ColorSpace",
	Olympus1RD2NoiseReduction:      "NoiseReduction",
	Olympus1RD2Engine:              "Engine",
	Olympus1RD2PictureMode:         "PictureMode",
	Olympus1RD2PMSaturation:        "PMSaturation",
	Olympus1RD2PMContrast:          "PMContrast",
	Olympus1RD2PMSharpness:         "PMSharpness",
	Olympus1RD2PM_BWFilter:         "PM_BWFilter",
	Olympus1RD2PMPictureTone:       "PMPictureTone",
	Olympus1RD2DevGradation:        "DevGradation",
	Olympus1RD2DevSaturation3:      "DevSaturation3",
	Olympus1RD2AutoGradation:       "AutoGradation",
	Olympus1RD2PMNoiseFilter:       "PMNoiseFilter",
	Olympus1RD2ArtFilter:           "ArtFilter",
}

// Tags in the Olympus1 Image Processing IFD
// ExifTool lib/Image/ExifTool/Olympus.pm ::ImageProcessing
const (
	Olympus1IPVersion               = 0x0
	Olympus1IPWB_RBLevels           = 0x100
	Olympus1IPWB_RBLevels3000K      = 0x102
	Olympus1IPWB_RBLevels3300K      = 0x103
	Olympus1IPWB_RBLevels3600K      = 0x104
	Olympus1IPWB_RBLevels3900K      = 0x105
	Olympus1IPWB_RBLevels4000K      = 0x106
	Olympus1IPWB_RBLevels4300K      = 0x107
	Olympus1IPWB_RBLevels4500K      = 0x108
	Olympus1IPWB_RBLevels4800K      = 0x109
	Olympus1IPWB_RBLevels5300K      = 0x10A
	Olympus1IPWB_RBLevels6000K      = 0x10B
	Olympus1IPWB_RBLevels6600K      = 0x10C
	Olympus1IPWB_RBLevels7500K      = 0x10D
	Olympus1IPWB_RBLevelsCWB1       = 0x10E
	Olympus1IPWB_RBLevelsCWB2       = 0x10F
	Olympus1IPWB_RBLevelsCWB3       = 0x110
	Olympus1IPWB_RBLevelsCWB4       = 0x111
	Olympus1IPWB_GLevel3000K        = 0x113
	Olympus1IPWB_GLevel3300K        = 0x114
	Olympus1IPWB_GLevel3600K        = 0x115
	Olympus1IPWB_GLevel3900K        = 0x116
	Olympus1IPWB_GLevel4000K        = 0x117
	Olympus1IPWB_GLevel4300K        = 0x118
	Olympus1IPWB_GLevel4500K        = 0x119
	Olympus1IPWB_GLevel4800K        = 0x11A
	Olympus1IPWB_GLevel5300K        = 0x11B
	Olympus1IPWB_GLevel6000K        = 0x11C
	Olympus1IPWB_GLevel6600K        = 0x11D
	Olympus1IPWB_GLevel7500K        = 0x11E
	Olympus1IPWB_GLevel             = 0x11F
	Olympus1IPColorMatrix           = 0x200
	Olympus1IPEnhancer              = 0x300
	Olympus1IPEnhancerValues        = 0x301
	Olympus1IPCoringFilter          = 0x310
	Olympus1IPCoringValues          = 0x311
	Olympus1IPBlackLevel2           = 0x600
	Olympus1IPGainBase              = 0x610
	Olympus1IPValidBits             = 0x611
	Olympus1IPCropLeft              = 0x612
	Olympus1IPCropTop               = 0x613
	Olympus1IPCropWidth             = 0x614
	Olympus1IPCropHeight            = 0x615
	Olympus1IPSensorCalibration     = 0x805
	Olympus1IPNoiseReduction2       = 0x1010
	Olympus1IPDistortionCorrection2 = 0x1011
	Olympus1IPShadingCompensation2  = 0x1012
	Olympus1IPMultipleExposureMode  = 0x101C
	Olympus1IPAspectRatio           = 0x1112
	Olympus1IPAspectFrame           = 0x1113
	Olympus1IPFacesDetected         = 0x1200
	Olympus1IPFaceDetectArea        = 0x1201
	Olympus1IPMaxFaces              = 0x1202
	Olympus1IPFaceDetectFrameSize   = 0x1203
	Olympus1IPFaceDetectFrameCrop   = 0x1207
	Olympus1IPCameraTemperature     = 0x1306
	Olympus1IPKeystoneCompensation  = 0x1900
	Olympus1IPKeystoneDirection     = 0x1901
	Olympus1IPKeystoneValue         = 0x1906
)

// Mapping from Olympus1 Image Processing tags to strings.
var Olympus1ImageProcessingTagNames = map[tiff.Tag]string{
	Olympus1IPVersion:               "Version",
	Olympus1IPWB_RBLevels:           "WB_RBLevels",
	Olympus1IPWB_RBLevels3000K:      "WB_RBLevels3000K",
	Olympus1IPWB_RBLevels3300K:      "WB_RBLevels3300K",
	Olympus1IPWB_RBLevels3600K:      "WB_RBLevels3600K",
	Olympus1IPWB_RBLevels3900K:      "WB_RBLevels3900K",
	Olympus1IPWB_RBLevels4000K:      "WB_RBLevels4000K",
	Olympus1IPWB_RBLevels4300K:      "WB_RBLevels4300K",
	Olympus1IPWB_RBLevels4500K:      "WB_RBLevels4500K",
	Olympus1IPWB_RBLevels4800K:      "WB_RBLevels4800K",
	Olympus1IPWB_RBLevels5300K:      "WB_RBLevels5300K",
	Olympus1IPWB_RBLevels6000K:      "WB_RBLevels6000K",
	Olympus1IPWB_RBLevels6600K:      "WB_RBLevels6600K",
	Olympus1IPWB_RBLevels7500K:      "WB_RBLevels7500K",
	Olympus1IPWB_RBLevelsCWB1:       "WB_RBLevelsCWB1",
	Olympus1IPWB_RBLevelsCWB2:       "WB_RBLevelsCWB2",
	Olympus1IPWB_RBLevelsCWB3:       "WB_RBLevelsCWB3",
	Olympus1IPWB_RBLevelsCWB4:       "WB_RBLevelsCWB4",
	Olympus1IPWB_GLevel3000K:        "WB_GLevel3000K",
	Olympus1IPWB_GLevel3300K:        "WB_GLevel3300K",
	Olympus1IPWB_GLevel3600K:        "WB_GLevel3600K",
	Olympus1IPWB_GLevel3900K:        "WB_GLevel3900K",
	Olympus1IPWB_GLevel4000K:        "WB_GLevel4000K",
	Olympus1IPWB_GLevel4300K:        "WB_GLevel4300K",
	Olympus1IPWB_GLevel4500K:        "WB_GLevel4500K",
	Olympus1IPWB_GLevel4800K:        "WB_GLevel4800K",
	Olympus1IPWB_GLevel5300K:        "WB_GLevel5300K",
	Olympus1IPWB_GLevel6000K:        "WB_GLevel6000K",
	Olympus1IPWB_GLevel6600K:        "WB_GLevel6600K",
	Olympus1IPWB_GLevel7500K:        "WB_GLevel7500K",
	Olympus1IPWB_GLevel:             "WB_GLevel",
	Olympus1IPColorMatrix:           "ColorMatrix",
	Olympus1IPEnhancer:              "Enhancer",
	Olympus1IPEnhancerValues:        "EnhancerValues",
	Olympus1IPCoringFilter:          "CoringFilter",
	Olympus1IPCoringValues:          "CoringValues",
	Olympus1IPBlackLevel2:           "BlackLevel2",
	Olympus1IPGainBase:              "GainBase",
	Olympus1IPValidBits:             "ValidBits",
	Olympus1IPCropLeft:              "CropLeft",
	Olympus1IPCropTop:               "CropTop",
	Olympus1IPCropWidth:             "CropWidth",
	Olympus1IPCropHeight:            "CropHeight",
	Olympus1IPSensorCalibration:     "SensorCalibration",
	Olympus1IPNoiseReduction2:       "NoiseReduction2",
	Olympus1IPDistortionCorrection2: "DistortionCorrection2",
	Olympus1IPShadingCompensation2:  "ShadingCompensation2",
	Olympus1IPMultipleExposureMode:  "MultipleExposureMode",
	Olympus1IPAspectRatio:           "AspectRatio",
	Olympus1IPAspectFrame:           "AspectFrame",
	Olympus1IPFacesDetected:         "FacesDetected",
	Olympus1IPFaceDetectArea:        "FaceDetectArea",
	Olympus1IPMaxFaces:              "MaxFaces",
	Olympus1IPFaceDetectFrameSize:   "FaceDetectFrameSize",
	Olympus1IPFaceDetectFrameCrop:   "FaceDetectFrameCrop",
	Olympus1IPCameraTemperature:     "CameraTemperature",
	Olympus1IPKeystoneCompensation:  "KeystoneCompensation",
	Olympus1IPKeystoneDirection:     "KeystoneDirection",
	Olympus1IPKeystoneValue:         "KeystoneValue",
}

// Tags in the Olympus1 Focus Info IFD
// ExifTool lib/Image/ExifTool/Olympus.pm ::FocusInfo
const (
	Olympus1FIVersion                  = 0x0
	Olympus1FIAutoFocus                = 0x209
	Olympus1FISceneDetect              = 0x210
	Olympus1FISceneArea                = 0x211
	Olympus1FISceneDetectData          = 0x212
	Olympus1FIZoomStepCount            = 0x300
	Olympus1FIFocusStepCount           = 0x301
	Olympus1FIFocusStepInfinity        = 0x303
	Olympus1FIFocusStepNear            = 0x304
	Olympus1FIFocusDistance            = 0x305
	Olympus1FIAFPoint                  = 0x308
	Olympus1FIAFInfo                   = 0x328
	Olympus1FIExternalFlash            = 0x1201
	Olympus1FIExternalFlashGuideNumber = 0x1203
	Olympus1FIExternalFlashBounce      = 0x1204
	Olympus1FIExternalFlashZoom        = 0x1205
	Olympus1FIInternalFlash            = 0x1208
	Olympus1FIManualFlash              = 0x1209
	Olympus1FIMacroLED                 = 0x120A
	Olympus1FISensorTemperature        = 0x1500
	Olympus1FIImageStabilization       = 0x1600
)

// Mapping from Olympus1 Focus Info tags to strings.
var Olympus1FocusInfoTagNames = map[tiff.Tag]string{
	Olympus1FIVersion:                  "Version",
	Olympus1FIAutoFocus:                "AutoFocus",
	Olympus1FISceneDetect:              "SceneDetect",
	Olympus1FISceneArea:                "SceneArea",
	Olympus1FISceneDetectData:          "SceneDetectData",
	Olympus1FIZoomStepCount:            "ZoomStepCount",
	Olympus1FIFocusStepCount:           "FocusStepCount",
	Olympus1FIFocusStepInfinity:        "FocusStepInfinity",
	Olympus1FIFocusStepNear:            "FocusStepNear",
	Olympus1FIFocusDistance:            "FocusDistance",
	Olympus1FIAFPoint:                  "AFPoint",
	Olympus1FIAFInfo:                   "AFInfo",
	Olympus1FIExternalFlash:            "ExternalFlash",
	Olympus1FIExternalFlashGuideNumber: "ExternalFlashGuideNumber",
	Olympus1FIExternalFlashBounce:      "ExternalFlashBounce",
	Olympus1FIExternalFlashZoom:        "ExternalFlashZoom",
	Olympus1FIInternalFlash:            "InternalFlash",
	Olympus1FIManualFlash:              "ManualFlash",
	Olympus1FIMacroLED:                 "MacroLED",
	Olympus1FISensorTemperature:        "SensorTemperature",
	Olympus1FIImageStabilization:       "ImageStabilization",
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
