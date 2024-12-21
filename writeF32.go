package gobuffer

import (
	"encoding/binary"
	"math"
)

func (b *GoBuffer) WriteF32LE(offset int64, data []float32) {
	if offset < 0 || (offset+int64(len(data))*4) > b.cap {
		panic(ErrBufferOverflow)
	}

	for i, val := range data {
		start := offset + int64(i*4)
		binary.LittleEndian.PutUint32(b.buf[start:], math.Float32bits(val))
	}
}

func (b *GoBuffer) WriteF32LENext(data []float32) {
	b.WriteF32LE(b.off, data)
	b.SeekByte(int64(len(data))*4, true)
}

func (b *GoBuffer) WriteF32BE(offset int64, data []float32) {
	if offset < 0 || (offset+int64(len(data))*4) > b.cap {
		panic(ErrBufferOverflow)
	}

	for i, val := range data {
		start := offset + int64(i*4)
		binary.BigEndian.PutUint32(b.buf[start:], math.Float32bits(val))
	}
}

func (b *GoBuffer) WriteF32BENext(data []float32) {
	b.WriteF32BE(b.off, data)
	b.SeekByte(int64(len(data))*4, true)
}
