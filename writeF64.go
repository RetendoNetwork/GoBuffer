package gobuffer

import (
	"encoding/binary"
	"math"
)

func (b *GoBuffer) WriteF64LE(off int64, data []float64) {
	if off < 0 || (off+int64(len(data))*8) > b.cap {
		panic(ErrBufferOverflow)
	}

	for i, val := range data {
		start := off + int64(i*8)
		binary.LittleEndian.PutUint64(b.buf[start:], math.Float64bits(val))
	}
}

func (b *GoBuffer) WriteF64LENext(data []float64) {
	b.WriteF64LE(b.off, data)
	b.SeekByte(int64(len(data))*8, true)
}

func (b *GoBuffer) WriteF64BE(off int64, data []float64) {
	if off < 0 || (off+int64(len(data))*8) > b.cap {
		panic(ErrBufferOverflow)
	}

	for i, val := range data {
		start := off + int64(i*8)
		binary.BigEndian.PutUint64(b.buf[start:], math.Float64bits(val))
	}
}

func (b *GoBuffer) WriteF64BENext(data []float64) {
	b.WriteF64BE(b.off, data)
	b.SeekByte(int64(len(data))*8, true)
}