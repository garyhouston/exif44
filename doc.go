/*
Package exif44 encodes and decodes Exif metadata in TIFF and JPEG
files. See the README in the repository for notes and limitations.

The high-level interfaces are Read / ReadFile for read-only processing
and ReadWrite / ReadWriteFile for read-write processing. These are
procedures that take a control structure with callbacks for
user-defined processing. Structuring it this way makes it possible for
the library itself to handle the details of processing JPEG files that
contain multiple images using the Multi-Picture Format (MPF), aka
Multi-Picture Object (MPO). This is a JPEG extension used by various
cameras to store multiple images in a single file, typically for large
preview images or stereoscopic images.

Errors that occur during decoding are passed to callbacks, and may
be encoded in a multierror structure; see
https://github.com/hashicorp/go-multierror.

The high-level interfaces are implemented using lower-level
interfaces, including the tiff66 and jpegsegs libraries which are
packaged separately.

*/
package exif44
