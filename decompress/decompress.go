package decompress

import "io"

type (
	Decompressor interface {
		Read(p []byte) (n int, err error)
	}

	// NewDecompressor defines the function that will be used to generate a new decompressor
	NewDecompressor func(io.Reader) (Decompressor, error)

)
