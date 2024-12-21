package gobuffer

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
