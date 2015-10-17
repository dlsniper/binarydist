// Package bzip2 implements the compression part missing from
// compress/bzip2 decompression by running bzip2 in another process.
package bzip2

import (
	"io"
	"os/exec"

	"github.com/dlsniper/binarydist/compress"
)

type bzip2Writer struct {
	c *exec.Cmd
	w io.WriteCloser
}

func (w *bzip2Writer) Write(b []byte) (int, error) {
	return w.w.Write(b)
}

func (w *bzip2Writer) Close() error {
	if err := w.w.Close(); err != nil {
		return err
	}
	return w.c.Wait()
}

// New creates a new bzip2 compressor
func New(w io.Writer) (compress.Compressor, error) {
	var (
		bw  = &bzip2Writer{}
		err error
	)

	bw.c = exec.Command("bzip2", "-c")
	bw.c.Stdout = w

	if bw.w, err = bw.c.StdinPipe(); err != nil {
		return nil, err
	}

	if err = bw.c.Start(); err != nil {
		return nil, err
	}

	return bw, nil
}
