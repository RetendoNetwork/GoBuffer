package gobuffer

func (b *GoBuffer) WriteI64LE(offset int64, data []int64) {
	if offset < 0 || (offset+int64(len(data))*8) > b.cap {
		panic(ErrBufferOverwrite)
	}

	for i := 0; i < len(data); i++ {
		start := offset + int64(i*8)
		b.buf[start] = byte(data[i])
		b.buf[start+1] = byte(data[i] >> 8)
		b.buf[start+2] = byte(data[i] >> 16)
		b.buf[start+3] = byte(data[i] >> 24)
		b.buf[start+4] = byte(data[i] >> 32)
		b.buf[start+5] = byte(data[i] >> 40)
		b.buf[start+6] = byte(data[i] >> 48)
		b.buf[start+7] = byte(data[i] >> 56)
	}
}

func (b *GoBuffer) WriteI64LENext(data []int64) {
	b.WriteI64LE(b.off, data)
	b.SeekByte(int64(len(data))*8, true)
}

func (b *GoBuffer) WriteI64BE(offset int64, data []int64) {
	if offset < 0 || (offset+int64(len(data))*8) > b.cap {
		panic(ErrBufferUnderwrite)
	}

	for i := 0; i < len(data); i++ {
		start := offset + int64(i*8)
		b.buf[start] = byte(data[i] >> 56)
		b.buf[start+1] = byte(data[i] >> 48)
		b.buf[start+2] = byte(data[i] >> 40)
		b.buf[start+3] = byte(data[i] >> 32)
		b.buf[start+4] = byte(data[i] >> 24)
		b.buf[start+5] = byte(data[i] >> 16)
		b.buf[start+6] = byte(data[i] >> 8)
		b.buf[start+7] = byte(data[i])
	}
}

func (b *GoBuffer) WriteI64BENext(data []int64) {
	b.WriteI64BE(b.off, data)
	b.SeekByte(int64(len(data))*8, true)
}
