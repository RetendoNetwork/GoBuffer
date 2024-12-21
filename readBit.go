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
