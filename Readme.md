# binarydist

[![GoDoc](https://godoc.org/github.com/dlsniper/binarydist?status.svg)](https://godoc.org/github.com/dlsniper/binarydist)
[![Circle CI](https://circleci.com/gh/dlsniper/binarydist.svg?style=svg&circle-token=abf9b5aa884649969ca6906d22ef61f8435c5375)](https://circleci.com/gh/dlsniper/binarydist)
[![Coverage Status](https://coveralls.io/repos/dlsniper/binarydist/badge.svg?branch=master&service=github)](https://coveralls.io/github/dlsniper/binarydist?branch=master)

Package binarydist implements binary diff and patch as described on
<http://www.daemonology.net/bsdiff/>. It reads and writes files
compatible with the tools there.

Documentation at <https://godoc.org/github.com/dlsniper/binarydist>.

The original implementation was created by <https://github.com/kr>.
It did not allowed for implementing a different compressor to enable
other use cases as such I had to change the implementation to support
this.

### TODO

- [ ] change the implementation from using package wide functions to
 structure wide functions
