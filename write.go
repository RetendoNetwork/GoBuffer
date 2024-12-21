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

func (b *GoBuffer) WriteByte(data byte) error {
	if b.off >= b.cap {
		return ErrBufferOverwrite
	}
	b.WriteBytes(b.off, []byte{data})
	b.SeekByte(1, true)
	return nil
}

func (b *GoBuffer) WriteByteNext(data byte) {
	b.WriteBytes(b.off, []byte{data})
	b.SeekByte(1, true)
}

func (b *GoBuffer) ClearBit(offset int64) {
	if offset >= b.cap || offset < 0 {
		err := ErrBufferUnderwrite
		if offset >= b.cap {
			err = ErrBufferOverwrite
		}
		panic(err)
	}

	byteIndex := offset / 8
	bitIndex := int(7 - (offset % 8))
	mask := byte(1 << uint(bitIndex))
	b.buf[byteIndex] &= ^mask

	b.Refresh()
}

// Use ClearBitAt() and not ClearBit() because you can maybe have an error with "signature".
func (b *GoBuffer) ClearBitAt(offset int64) {
	b.ClearBit(offset)
}

func (b *GoBuffer) ClearBitNext() {
	b.ClearBit(b.boff)

	b.SeekBit(1, true)
}
