# exif44
exif44 is a Go library for encoding and decoding Exif data in TIFF and JPEG files. It can be used to extract or edit metadata, but doesn't include functionality for processing images.

For documentation, see https://godoc.org/github.com/garyhouston/exif44.

## Notes and limitations
This library is still under construction and may change at any moment without backwards compatibility.

Since Exif stores metadata in TIFF format, this library makes use of the [tiff66 library](https://github.com/garyhouston/tiff66). The notes that apply to that library are relevant here. The [jpegsegs library](https://github.com/garyhouston/jpegsegs) can be used to decode a JPEG file into segments.

The exif44print program prints the IFDs (image file directories) and fields, either from a TIFF file or from the Exif segment of a JPEG file.

The exif44repack program decodes a TIFF file, or the Exif segment of a JPEG file, re-encodes it and writes it to a new file.

Metadata in JPEG files can also be stored in other formats such as XMP, which is not supported by this library. Both formats can be present in the same file.

As per tiff66, most maker note fields found in Exif are not currently decoded. In some cases they contain pointers which will be broken if the file is repacked.

This library makes no provision for modification of data in multiple threads. Mutexes etc., should be used as required.

Invalid files may result in vague errors or strange data structures.

'44' is an arbitrary number to distinguish this library from all the other Exif libraries.
