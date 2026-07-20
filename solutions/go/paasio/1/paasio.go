package paasio

import (
	"io"
	"sync"
)

// Reader holds read counts and underlying io.Reader
type reader struct {
	sync.RWMutex
	r    io.Reader
	n    int64
	nops int
}

// Writer holds read counts and underlying io.Writer
type writer struct {
	sync.RWMutex
	w    io.Writer
	n    int64
	nops int
}

// ReaderWriter holds Reader and Writer
type readerWriter struct {
	ReadCounter
	WriteCounter
}

// NewReadCounter returns new ReadCounter
func NewReadCounter(r io.Reader) ReadCounter {
	return &reader{
		r: r,
	}
}

// NewWriteCounter returns new WriteCounter
func NewWriteCounter(w io.Writer) WriteCounter {
	return &writer{
		w: w,
	}
}

// NewReadWriteCounter returns a new ReadWriteCounter
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &readerWriter{
		NewReadCounter(rw),
		NewWriteCounter(rw),
	}
}

func (r *reader) Read(p []byte) (int, error) {
	r.Lock()
	defer r.Unlock()

	n, err := r.r.Read(p)
	r.n += int64(n)
	r.nops++
	return n, err
}

// ReadCount return number of bytes read and no of read operations
func (r *reader) ReadCount() (n int64, nops int) {
	r.RLock()
	defer r.RUnlock()

	return r.n, r.nops
}

func (w *writer) Write(p []byte) (int, error) {
	w.Lock()
	defer w.Unlock()

	n, err := w.w.Write(p)
	w.n += int64(n)
	w.nops++
	return n, err
}

// WriteCount return number of bytes written and no of write operations
func (w *writer) WriteCount() (n int64, nops int) {
	w.RLock()
	defer w.RUnlock()

	return w.n, w.nops
}
