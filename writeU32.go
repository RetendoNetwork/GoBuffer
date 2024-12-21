package gobuffer

// WriteU32LE writes a slice of uint32 to the buffer at the specified offset in little-endian order.
func (b *GoBuffer) WriteU32LE(off int64, data []uint32) {
	if off < 0 || (off+int64(len(data))*4) > b.cap {
		panic(ErrBufferOverwrite)
	}

	for i, val := range data {
		start := off + int64(i*4)
		b.buf[start] = byte(val)
		b.buf[start+1] = byte(val >> 8)
		b.buf[start+2] = byte(val >> 16)
		b.buf[start+3] = byte(val >> 24)
	}
}

// WriteU32LENext writes a slice of uint32 to the buffer at the current offset in little-endian order.
func (b *GoBuffer) WriteU32LENext(data []uint32) {
	b.WriteU32LE(b.off, data)
	b.SeekByte(int64(len(data))*4, true)
}

// WriteU32BE writes a slice of uint32 to the buffer at the specified offset in big-endian order.
func (b *GoBuffer) WriteU32BE(off int64, data []uint32) {
	if off < 0 || (off+int64(len(data))*4) > b.cap {
		panic(ErrBufferOverwrite)
	}

	for i, val := range data {
		start := off + int64(i*4)
		b.buf[start] = byte(val >> 24)
		b.buf[start+1] = byte(val >> 16)
		b.buf[start+2] = byte(val >> 8)
		b.buf[start+3] = byte(val)
	}
}

// WriteU32BENext writes a slice of uint32 to the buffer at the current offset in big-endian order.
func (b *GoBuffer) WriteU32BENext(data []uint32) {
	b.WriteU32BE(b.off, data)
	b.SeekByte(int64(len(data))*4, true)
}
