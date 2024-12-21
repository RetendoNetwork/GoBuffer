package gobuffer

import (
	"fmt"
)

type GoBuffer struct {
	buf  []byte
	off  int64
	cap  int64
	boff int64
	bcap int64
}

func NewGoBuffer(slices ...[]byte) *GoBuffer {
	buf := &GoBuffer{
		buf:  []byte{},
		off:  0,
		boff: 0,
	}

	if len(slices) == 0 {
		buf.Refresh()
		return buf
	}

	if len(slices) == 1 {
		buf.buf = slices[0]
		buf.Refresh()
		return buf
	}

	for _, slice := range slices {
		buf.buf = append(buf.buf, slice...)
	}

	buf.Refresh()
	return buf
}

func (b *GoBuffer) SeekBit(offset int64, relative bool) {
	switch relative {
	case true:
		b.boff = b.boff + offset
	default:
		b.boff = offset
	}
}

func (b *GoBuffer) FlipBit(offset int64) {
	if offset < 0 || offset >= b.bcap {
		panic(fmt.Errorf("invalid offset %d, out of bounds", offset))
	}

	byteIndex := offset / 8
	bitIndex := 7 - (offset % 8)

	b.buf[byteIndex] ^= 1 << uint(bitIndex)
}

func (b *GoBuffer) SeekByte(offset int64, relative bool) {
	if relative {
		b.off += offset
	} else {
		b.off = offset
	}
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

func (b *GoBuffer) AlignBit() {
	b.boff = b.off << 3
}

func divMod(a, b int64) (int64, int64) {
	return a / b, a % b
}

func (b *GoBuffer) AlignByte() {
	quotient, _ := divMod(b.boff, 8)

	b.off = quotient
}

func (b *GoBuffer) Refresh() {
	b.bcap = int64(len(b.buf))

	b.cap = b.bcap * 8
}

// Original code of Grow from https://github.com/habak67/gobuffer/blob/master/buffer.go#L151
func (b *GoBuffer) Grow(size int64) {
	if size < 0 {
		panic(fmt.Errorf("invalid size: cannot be negative"))
	}

	if size <= b.bcap {
		b.buf = b.buf[0 : b.off+size]
		b.Refresh()
		return
	}

	newCapacity := b.bcap * 2
	if newCapacity < size {
		newCapacity = size
	}

	tmp := make([]byte, newCapacity)
	copy(tmp, b.buf)

	b.buf = tmp
	b.Refresh()
}

func (b *GoBuffer) Bytes() []byte {
	return b.buf
}

func (b *GoBuffer) ByteCapacity() int64 {
	return b.cap
}

func (b *GoBuffer) BitCapacity() int64 {
	return b.bcap
}

func (b *GoBuffer) ByteOffset() int64 {
	return b.off
}

func (b *GoBuffer) BitOffset() int64 {
	return b.boff
}
