package gobuffer

func (b *GoBuffer) WriteBytes(offset int64, data []byte) {
	if offset < 0 {
		panic(ErrBufferUnderwrite)
	}

	if offset+int64(len(data)) > b.cap {
		panic(ErrBufferOverwrite)
	}

	copy(b.buf[offset:], data)
}

func (b *GoBuffer) WriteBytesNext(data []byte) {
	b.WriteBytes(b.off, data)
	b.SeekByte(int64(len(data)), true)
}
