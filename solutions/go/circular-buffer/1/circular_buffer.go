// Package circular implements simple routines for circular bufffer
package circular

import "errors"

// Buffer holds items and current read index
type Buffer struct {
	data      []byte
	count     int
	readIndex int
}

// NewBuffer creates and returns a new Buffer
func NewBuffer(size int) *Buffer {
	return &Buffer{
		data: make([]byte, size),
	}
}

// ReadByte returns byte from current read postion
func (b *Buffer) ReadByte() (byte, error) {
	if b.count == 0 {
		return 0, errors.New("cannot read empty buffer")
	}

	ri := b.readIndex % len(b.data)
	v := b.data[ri]
	b.readIndex = ri + 1
	b.count--

	return v, nil
}

// WriteByte appends byte to buffer returns error if buffer is full
func (b *Buffer) WriteByte(v byte) error {
	if b.count == len(b.data) {
		return errors.New("cannot write to full buffer")
	}

	wi := (b.readIndex + b.count) % len(b.data)
	b.data[wi] = v
	b.count++

	return nil
}

// Overwrite writes value to buffer, if full overwrites the first value
func (b *Buffer) Overwrite(v byte) {
	if b.count == len(b.data) {
		b.ReadByte()
	}
	b.WriteByte(v)
}

// Reset resets buffer
func (b *Buffer) Reset() {
	b.count = 0
}
