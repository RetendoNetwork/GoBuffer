package gobuffer

import (
	"fmt"
)

func (b *GoBuffer) ReadBit(out *byte, offset int64) error {
	byteIndex := offset / 8
	bitIndex := 7 - (offset % 8)

	if byteIndex >= int64(len(b.buf)) {
		return fmt.Errorf("out of bounds")
	}

	*out = (b.buf[byteIndex] >> uint(bitIndex)) & 1
	return nil
}

func (b *GoBuffer) ReadBitNext() byte {
	var out byte
	err := b.ReadBit(&out, b.boff)
	if err != nil {
		panic(err)
	}
	b.SeekBit(1, true)
	return out
}

func (b *GoBuffer) ReadBits(out *uint64, off, n int64) error {
	var result uint64
	var bout byte

	for i := int64(0); i < n; i++ {
		err := b.ReadBit(&bout, off+i)
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
