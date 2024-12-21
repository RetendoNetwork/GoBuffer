package gobuffer

func (b *GoBuffer) WriteI16LE(off int64, data []int16) {
	if off < 0 || (off+int64(len(data))*2) > b.cap {
		panic(ErrBufferOverflow)
	}

	for i, val := range data {
		start := off + int64(i*2)
		b.buf[start] = byte(val)
		b.buf[start+1] = byte(val >> 8)
	}
}

func (b *GoBuffer) WriteI16LENext(data []int16) {
	b.WriteI16LE(b.off, data)
	b.SeekByte(int64(len(data))*2, true)
}

func (b *GoBuffer) WriteI16BE(off int64, data []int16) {
	if off < 0 || (off+int64(len(data))*2) > b.cap {
		panic(ErrBufferOverflow)
	}

	for i, val := range data {
		start := off + int64(i*2)
		b.buf[start] = byte(val >> 8)
		b.buf[start+1] = byte(val)
	}
}

func (b *GoBuffer) WriteI16BENext(data []int16) {
	b.WriteI16BE(b.off, data)
	b.SeekByte(int64(len(data))*2, true)
}
