# binarydist

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
