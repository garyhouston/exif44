package exif44

import (
	tiff "github.com/garyhouston/tiff66"
)

// Tag names are from ExifTool 10.49, which also has information about
// how to interpret the values.

// Tags in the "Panasonic 1" maker note.
// ExifTool lib/Image/ExifTool/Panasonic.pm
const (
	Pan1ImageQuality               = 0x1
	Pan1FirmwareVersion            = 0x2
	Pan1WhiteBalance               = 0x3
	Pan1FocusMode                  = 0x7
	Pan1AFAreaMode                 = 0xF
	Pan1ImageStabilization         = 0x1A
	Pan1MacroMode                  = 0x1C
	Pan1ShootingMode               = 0x1F
	Pan1Audio                      = 0x20
	Pan1DataDump                   = 0x21
	Pan1WhiteBalanceBias           = 0x23
	Pan1FlashBias                  = 0x24
	Pan1InternalSerialNumber       = 0x25
	Pan1PanasonicExifVersion       = 0x26
	Pan1ColorEffect                = 0x28
	Pan1TimeSincePowerOn           = 0x29
	Pan1BurstMode                  = 0x2A
	Pan1SequenceNumber             = 0x2B
	Pan1ContrastMode               = 0x2C
	Pan1NoiseReduction             = 0x2D
	Pan1SelfTimer                  = 0x2E
	Pan1Rotation                   = 0x30
	Pan1AFAssistLamp               = 0x31
	Pan1ColorMode                  = 0x32
	Pan1BabyAge1                   = 0x33
	Pan1OpticalZoomMode            = 0x34
	Pan1ConversionLens             = 0x35
	Pan1TravelDay                  = 0x36
	Pan1Contrast                   = 0x39
	Pan1WorldTimeLocation          = 0x3A
	Pan1TextStamp1                 = 0x3B
	Pan1ProgramISO                 = 0x3C
	Pan1AdvancedSceneType          = 0x3D
	Pan1TextStamp2                 = 0x3E
	Pan1FacesDetected              = 0x3F
	Pan1Saturation                 = 0x40
	Pan1Sharpness                  = 0x41
	Pan1FilmMode                   = 0x42
	Pan1ColorTempKelvin            = 0x44
	Pan1BracketSettings            = 0x45
	Pan1WBShiftAB                  = 0x46
	Pan1WBShiftGM                  = 0x47
	Pan1FlashCurtain               = 0x48
	Pan1LongExposureNoiseReduction = 0x49
	Pan1PanasonicImageWidth        = 0x4B
	Pan1PanasonicImageHeight       = 0x4C
	Pan1AFPointPosition            = 0x4D
	Pan1FaceDetInfo                = 0x4E
	Pan1LensType                   = 0x51
	Pan1LensSerialNumber           = 0x52
	Pan1AccessoryType              = 0x53
	Pan1AccessorySerialNumber      = 0x54
	Pan1Transform1                 = 0x59
	Pan1IntelligentExposure        = 0x5D
	Pan1LensFirmwareVersion        = 0x60
	Pan1FaceRecInfo                = 0x61
	Pan1FlashWarning               = 0x62
	Pan1RecognizedFaceFlags        = 0x63
	Pan1Title                      = 0x65
	Pan1BabyName                   = 0x66
	Pan1Location                   = 0x67
	Pan1Country                    = 0x69
	Pan1State                      = 0x6B
	Pan1City1                      = 0x6D
	Pan1Landmark                   = 0x6F
	Pan1IntelligentResolution      = 0x70
	Pan1BurstSpeed                 = 0x77
	Pan1IntelligentDRange          = 0x79
	Pan1ClearRetouch               = 0x7C
	Pan1City2                      = 0x80
	Pan1ManometerPressure          = 0x86
	Pan1PhotoStyle                 = 0x89
	Pan1ShadingCompensation        = 0x8A
	Pan1AccelerometerZ             = 0x8C
	Pan1AccelerometerX             = 0x8D
	Pan1AccelerometerY             = 0x8E
	Pan1CameraOrientation          = 0x8F
	Pan1RollAngle                  = 0x90
	Pan1PitchAngle                 = 0x91
	Pan1SweepPanoramaDirection     = 0x93
	Pan1SweepPanoramaFieldOfView   = 0x94
	Pan1TimerRecording             = 0x96
	Pan1InternalNDFilter           = 0x9D
	Pan1HDR                        = 0x9E
	Pan1ShutterType                = 0x9F
	Pan1ClearRetouchValue          = 0xA3
	Pan1TouchAE                    = 0xAB
	Pan1TimeStamp                  = 0xAF
	Pan1PrintIM                    = 0xE00
	Pan1MakerNoteVersion           = 0x8000
	Pan1SceneMode                  = 0x8001
	Pan1WBRedLevel                 = 0x8004
	Pan1WBGreenLevel               = 0x8005
	Pan1WBBlueLevel                = 0x8006
	Pan1FlashFired                 = 0x8007
	Pan1TextStamp3                 = 0x8008
	Pan1TextStamp4                 = 0x8009
	Pan1BabyAge2                   = 0x8010
	Pan1Transform2                 = 0x8012
)

// Mapping from Panasonic 1 tags to strings.
var Panasonic1TagNames = map[tiff.Tag]string{
	Pan1ImageQuality:               "ImageQuality",
	Pan1FirmwareVersion:            "FirmwareVersion",
	Pan1WhiteBalance:               "WhiteBalance",
	Pan1FocusMode:                  "FocusMode",
	Pan1AFAreaMode:                 "AFAreaMode",
	Pan1ImageStabilization:         "ImageStabilization",
	Pan1MacroMode:                  "MacroMode",
	Pan1ShootingMode:               "ShootingMode",
	Pan1Audio:                      "Audio",
	Pan1DataDump:                   "DataDump",
	Pan1WhiteBalanceBias:           "WhiteBalanceBias",
	Pan1FlashBias:                  "FlashBias",
	Pan1InternalSerialNumber:       "InternalSerialNumber",
	Pan1PanasonicExifVersion:       "PanasonicExifVersion",
	Pan1ColorEffect:                "ColorEffect",
	Pan1TimeSincePowerOn:           "TimeSincePowerOn",
	Pan1BurstMode:                  "BurstMode",
	Pan1SequenceNumber:             "SequenceNumber",
	Pan1ContrastMode:               "ContrastMode",
	Pan1NoiseReduction:             "NoiseReduction",
	Pan1SelfTimer:                  "SelfTimer",
	Pan1Rotation:                   "Rotation",
	Pan1AFAssistLamp:               "AFAssistLamp",
	Pan1ColorMode:                  "ColorMode",
	Pan1BabyAge1:                   "BabyAge1",
	Pan1OpticalZoomMode:            "OpticalZoomMode",
	Pan1ConversionLens:             "ConversionLens",
	Pan1TravelDay:                  "TravelDay",
	Pan1Contrast:                   "Contrast",
	Pan1WorldTimeLocation:          "WorldTimeLocation",
	Pan1TextStamp1:                 "TextStamp1",
	Pan1ProgramISO:                 "ProgramISO",
	Pan1AdvancedSceneType:          "AdvancedSceneType",
	Pan1TextStamp2:                 "TextStamp2",
	Pan1FacesDetected:              "FacesDetected",
	Pan1Saturation:                 "Saturation",
	Pan1Sharpness:                  "Sharpness",
	Pan1FilmMode:                   "FilmMode",
	Pan1ColorTempKelvin:            "ColorTempKelvin",
	Pan1BracketSettings:            "BracketSettings",
	Pan1WBShiftAB:                  "WBShiftAB",
	Pan1WBShiftGM:                  "WBShiftGM",
	Pan1FlashCurtain:               "FlashCurtain",
	Pan1LongExposureNoiseReduction: "LongExposureNoiseReduction",
	Pan1PanasonicImageWidth:        "PanasonicImageWidth",
	Pan1PanasonicImageHeight:       "PanasonicImageHeight",
	Pan1AFPointPosition:            "AFPointPosition",
	Pan1FaceDetInfo:                "FaceDetInfo",
	Pan1LensType:                   "LensType",
	Pan1LensSerialNumber:           "LensSerialNumber",
	Pan1AccessoryType:              "AccessoryType",
	Pan1AccessorySerialNumber:      "AccessorySerialNumber",
	Pan1Transform1:                 "Transform1",
	Pan1IntelligentExposure:        "IntelligentExposure",
	Pan1LensFirmwareVersion:        "LensFirmwareVersion",
	Pan1FaceRecInfo:                "FaceRecInfo",
	Pan1FlashWarning:               "FlashWarning",
	Pan1RecognizedFaceFlags:        "RecognizedFaceFlags",
	Pan1Title:                      "Title",
	Pan1BabyName:                   "BabyName",
	Pan1Location:                   "Location",
	Pan1Country:                    "Country",
	Pan1State:                      "State",
	Pan1City1:                      "City1",
	Pan1Landmark:                   "Landmark",
	Pan1IntelligentResolution:      "IntelligentResolution",
	Pan1BurstSpeed:                 "BurstSpeed",
	Pan1IntelligentDRange:          "IntelligentD-Range",
	Pan1ClearRetouch:               "ClearRetouch",
	Pan1City2:                      "City2",
	Pan1ManometerPressure:          "ManometerPressure",
	Pan1PhotoStyle:                 "PhotoStyle",
	Pan1ShadingCompensation:        "ShadingCompensation",
	Pan1AccelerometerZ:             "AccelerometerZ",
	Pan1AccelerometerX:             "AccelerometerX",
	Pan1AccelerometerY:             "AccelerometerY",
	Pan1CameraOrientation:          "CameraOrientation",
	Pan1RollAngle:                  "RollAngle",
	Pan1PitchAngle:                 "PitchAngle",
	Pan1SweepPanoramaDirection:     "SweepPanoramaDirection",
	Pan1SweepPanoramaFieldOfView:   "SweepPanoramaFieldOfView",
	Pan1TimerRecording:             "TimerRecording",
	Pan1InternalNDFilter:           "InternalNDFilter",
	Pan1HDR:                        "HDR",
	Pan1ShutterType:                "ShutterType",
	Pan1ClearRetouchValue:          "ClearRetouchValue",
	Pan1TouchAE:                    "TouchAE",
	Pan1TimeStamp:                  "TimeStamp",
	Pan1PrintIM:                    "PrintIM",
	Pan1MakerNoteVersion:           "MakerNoteVersion",
	Pan1SceneMode:                  "SceneMode",
	Pan1WBRedLevel:                 "WBRedLevel",
	Pan1WBGreenLevel:               "WBGreenLevel",
	Pan1WBBlueLevel:                "WBBlueLevel",
	Pan1FlashFired:                 "FlashFired",
	Pan1TextStamp3:                 "TextStamp3",
	Pan1TextStamp4:                 "TextStamp4",
	Pan1BabyAge2:                   "BabyAge2",
	Pan1Transform2:                 "Transform2",
}
