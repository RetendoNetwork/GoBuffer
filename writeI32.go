package gobuffer

func (b *GoBuffer) WriteI32LE(off int64, data []int32) {
	if off < 0 || (off+int64(len(data))*4) > b.cap {
		panic(ErrBufferOverwrite)
	}

	for i, value := range data {
		start := off + int64(i*4)
		b.buf[start] = byte(value)
		b.buf[start+1] = byte(value >> 8)
		b.buf[start+2] = byte(value >> 16)
		b.buf[start+3] = byte(value >> 24)
	}
}

func (b *GoBuffer) WriteI32LENext(data []int32) {
	b.WriteI32LE(b.off, data)
	b.SeekByte(int64(len(data))*4, true)
}

func (b *GoBuffer) WriteI32BE(off int64, data []int32) {
	if off < 0 || (off+int64(len(data))*4) > b.cap {
		panic(ErrBufferUnderwrite)
	}

	for i, value := range data {
		start := off + int64(i*4)
		b.buf[start] = byte(value >> 24)
		b.buf[start+1] = byte(value >> 16)
		b.buf[start+2] = byte(value >> 8)
		b.buf[start+3] = byte(value)
	}
}

func (b *GoBuffer) WriteI32BENext(data []int32) {
	b.WriteI32BE(b.off, data)
	b.SeekByte(int64(len(data))*4, true)
}
