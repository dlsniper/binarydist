package bzip2

import (
	"compress/bzip2"
	"io"

	"github.com/dlsniper/binarydist/decompress"
)

type bzip2Reader struct {
	r io.Reader
}

func (br *bzip2Reader) Read(p []byte) (n int, err error) {
	return br.r.Read(p)
}

func New(r io.Reader) (decompress.Decompressor, error) {
	return &bzip2Reader{
		r: bzip2.NewReader(r),
	}, nil
}
