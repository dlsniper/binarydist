package binarydist

import (
	"io"

	"github.com/dlsniper/binarydist/compress"
	"github.com/dlsniper/binarydist/compress/bzip2"
)

var (
	compressorFactory = bzip2.New
)

// SetCompressor will set the current used compressor
func SetCompressor(c compress.NewCompressor) {
	compressorFactory = c
}

func newCompressor(w io.Writer) (compress.Compressor, error) {
	return compressorFactory(w)
}
