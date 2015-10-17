package decompress

import "io"

type (
	Decompressor interface {
		Read([]byte) (int, error)
		Close() (error)
	}

	// NewDecompressor defines the function that will be used to generate a new decompressor
	NewDecompressor func(io.Reader) (Decompressor, error)

)
