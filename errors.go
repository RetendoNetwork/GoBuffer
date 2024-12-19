package gobuffer

import "errors"

var (
	ErrBufferOverwrite  = errors.New("buffer overwrite error")
	ErrBufferUnderwrite = errors.New("buffer underwrite error")
)
