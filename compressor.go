package binarydist

import (
	"io"

	"github.com/dlsniper/binarydist/compress"
	comBzip2 "github.com/dlsniper/binarydist/compress/bzip2"
	"github.com/dlsniper/binarydist/decompress"
	decompBzip2 "github.com/dlsniper/binarydist/decompress/bzip2"
)

var (
	compressorFactory   = comBzip2.New
	decompressorFactory = decompBzip2.New
)

// SetCompressor will set the current used compressor
func SetCompressor(c compress.NewCompressor) {
	compressorFactory = c
}

// SetDecompressor will set the current used compressor
func SetDecompressor(d decompress.NewDecompressor) {
	decompressorFactory = d
}

func newCompressor(w io.Writer) (compress.Compressor, error) {
	return compressorFactory(w)
}

func newDecompressor(r io.Reader) (decompress.Decompressor, error) {
	return decompressorFactory(r)
}
