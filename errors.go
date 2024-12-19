package gobuffer

import "errors"

var (
	BufferOverwriteError  = errors.New("buffer overwrite error")
	BufferUnderwriteError = errors.New("buffer underwrite error")
)
