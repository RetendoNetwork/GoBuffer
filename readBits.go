package gobuffer

func (b *GoBuffer) ReadBits(out *uint64, offset, n int64) error {
	var result uint64
	var bout byte

	for i := int64(0); i < n; i++ {
		err := b.ReadBit(&bout, offset+i)
		if err != nil {
			return err
		}

		result = (result << 1) | uint64(bout)
	}

	*out = result
	return nil
}

func (b *GoBuffer) ReadBitsNext(n int64) uint64 {
	var out uint64
	err := b.ReadBits(&out, b.boff, n)
	if err != nil {
		panic(err)
	}
	b.SeekBit(n, true)
	return out
}
