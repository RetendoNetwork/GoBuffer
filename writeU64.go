package gobuffer

// WriteU64LE writes a slice of uint64 to the buffer at the specified offset in little-endian order.
func (b *GoBuffer) WriteU64LE(offset int64, data []uint64) {
	if offset < 0 || (offset+int64(len(data))*8) > b.cap {
		panic(ErrBufferOverwrite)
	}

	for i, val := range data {
		start := offset + int64(i*8)
		b.buf[start] = byte(val)
		b.buf[start+1] = byte(val >> 8)
		b.buf[start+2] = byte(val >> 16)
		b.buf[start+3] = byte(val >> 24)
		b.buf[start+4] = byte(val >> 32)
		b.buf[start+5] = byte(val >> 40)
		b.buf[start+6] = byte(val >> 48)
		b.buf[start+7] = byte(val >> 56)
	}
}

// WriteU64LENext writes a slice of uint64 to the buffer at the current offset in little-endian order.
func (b *GoBuffer) WriteU64LENext(data []uint64) {
	b.WriteU64LE(b.off, data)
	b.SeekByte(int64(len(data))*8, true)
}

// WriteU64BE writes a slice of uint64 to the buffer at the specified offset in big-endian order.
func (b *GoBuffer) WriteU64BE(offset int64, data []uint64) {
	if offset < 0 || (offset+int64(len(data))*8) > b.cap {
		panic(ErrBufferOverwrite)
	}

	for i, val := range data {
		start := offset + int64(i*8)
		b.buf[start] = byte(val >> 56)
		b.buf[start+1] = byte(val >> 48)
		b.buf[start+2] = byte(val >> 40)
		b.buf[start+3] = byte(val >> 32)
		b.buf[start+4] = byte(val >> 24)
		b.buf[start+5] = byte(val >> 16)
		b.buf[start+6] = byte(val >> 8)
		b.buf[start+7] = byte(val)
	}
}

// WriteU64BENext writes a slice of uint64 to the buffer at the current offset in big-endian order.
func (b *GoBuffer) WriteU64BENext(data []uint64) {
	b.WriteU64BE(b.off, data)
	b.SeekByte(int64(len(data))*8, true)
}
