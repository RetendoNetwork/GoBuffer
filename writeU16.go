package gobuffer

func (b *GoBuffer) WriteU16LE(off int64, data []uint16) {
	if off < 0 || (off+int64(len(data))*2) > b.cap {
		panic(ErrBufferOverwrite)
	}

	for i, val := range data {
		start := off + int64(i*2)
		b.buf[start] = byte(val)
		b.buf[start+1] = byte(val >> 8)
	}
}

func (b *GoBuffer) WriteU16LENext(data []uint16) {
	b.WriteU16LE(b.off, data)
	b.SeekByte(int64(len(data))*2, true)
}

func (b *GoBuffer) WriteU16BE(off int64, data []uint16) {
	if off < 0 || (off+int64(len(data))*2) > b.cap {
		panic(ErrBufferUnderwrite)
	}

	for i, val := range data {
		start := off + int64(i*2)
		b.buf[start] = byte(val >> 8)
		b.buf[start+1] = byte(val)
	}
}

func (b *GoBuffer) WriteU16BENext(data []uint16) {
	b.WriteU16BE(b.off, data)
	b.SeekByte(int64(len(data))*2, true)
}
