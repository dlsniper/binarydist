package compress

import "io"

type (
	// Compressor defines the functions that a compressor must have in order
	// to describe a correct compressing functionality
	Compressor interface {
		Write(b []byte) (int, error)
		Close() error
	}

	// NewCompressor defines the function that will be used to generate a new compressor
	NewCompressor func(io.Writer) (Compressor, error)
)
