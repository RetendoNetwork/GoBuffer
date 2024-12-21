package gobuffer

import "errors"

var (
	ErrBufferOverwrite        = errors.New("buffer overwrite error")
	ErrBufferUnderwrite       = errors.New("buffer underwrite error")
	ErrBufferInvalidByteCount = errors.New("invalid byte count error")
	ErrBufferNegativeRead     = errors.New("negative read error")
	ErrBufferOverflow         = errors.New("buffer overflow error")
)
