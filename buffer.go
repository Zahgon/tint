package tint

import "sync"

type buffer []byte

var bufPool = sync.Pool{
	New: func() any {
		b := make(buffer, 0, 1024)
		return (*buffer)(&b)
	},
}

func newBuffer() *buffer { _ = "STUB: not implemented"; return nil }

func (b *buffer) Free() {
	_ = "STUB: not implemented"
	// To reduce peak allocation, return only smaller buffers to the pool.
	return
}

func (b *buffer) Write(bytes []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *buffer) WriteByte(char byte) error { _ = "STUB: not implemented"; return nil }

func (b *buffer) WriteString(str string) (int, error) { _ = "STUB: not implemented"; return 0, nil }
