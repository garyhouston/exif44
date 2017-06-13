package exif44

import (
	tiff "github.com/garyhouston/tiff66"
)

// Tag names are from ExifTool 10.49, which also has information about
// how to interpret the values.

// Tags in the "Panasonic 1" maker note.
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

// Mapping from Panasonic 1 tags to strings.
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
