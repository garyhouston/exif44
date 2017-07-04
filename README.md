# exif44
exif44 is a Go library for encoding and decoding Exif data in TIFF and JPEG files. It can be used to extract or edit metadata, but doesn't include functionality for processing images.

For documentation, see https://godoc.org/github.com/garyhouston/exif44.

## Notes and limitations
This library is still under construction and may change at any moment without backwards compatibility.

Since Exif stores metadata in TIFF format, this library makes use of the [tiff66 library](https://github.com/garyhouston/tiff66). The notes that apply to that library are relevant here. The [jpegsegs library](https://github.com/garyhouston/jpegsegs) can be used to decode a JPEG file into segments.

The exif44print program prints the IFDs (image file directories) and fields, either from a TIFF file or from the Exif segment of a JPEG file.

The exif44repack program decodes a TIFF file, or the Exif segment of a JPEG file, re-encodes it and writes it to a new file.

The exif44offsets program displays the byte layout of IFDs. If the TIFF data doesn't start at the beginning of the file, it will print a header with the offset, and remaining positions will be relative to the start of the TIFF header. Each line of output will have a start and ending position and a description of where the data originated. If the description is blank, nothing is known and it may a data region for a field not supported by the tiff66 library. Unused single bytes are likely to be fillers to align the next IFD. Maker notes are included in IFD external data.

The exif44addloc program adds location coordinates (GPS) to a JPEG or TIFF file. It's run as 'exif44addloc latitude longitude file-in file-out', with the coordinates expressed as decimal numbers.

Metadata in JPEG files can also be stored in other formats such as XMP, which is not supported by this library. Both formats can be present in the same file.

As per tiff66, most maker note fields found in Exif are not currently decoded. In some cases they contain pointers which will be broken if a file is rewritten by this library.

This library makes no provision for modification of data in multiple threads. Mutexes etc., should be used as required.

'44' is an arbitrary number to distinguish this library from all the other Exif libraries.
