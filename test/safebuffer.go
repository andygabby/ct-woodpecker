package test

import (
	"bytes"
	"sync"
)

// SafeBuffer is a wrapper around a bytes.Buffer that is made safe for
// concurrent access. This is required for making a logger backed by a bytes
// buffer that can be used by a woodpecker instance across multiple goroutines
// without a data race.
type SafeBuffer struct {
	b bytes.Buffer
	m sync.RWMutex
}

// Reset safely resets the wrapped buffer
func (b *SafeBuffer) Reset() {
	b.m.Lock()
	defer b.m.Unlock()
	b.b.Reset()
}

// Read safely reads from the wrapped buffer
func (b *SafeBuffer) Read(p []byte) (n int, err error) {
	b.m.RLock()
	defer b.m.RUnlock()
	return b.b.Read(p)
}

// Write safely writes to the wrapped buffer
func (b *SafeBuffer) Write(p []byte) (n int, err error) {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.Write(p)
}

// String safely reads from the wrapped buffer
func (b *SafeBuffer) String() string {
	b.m.RLock()
	defer b.m.RUnlock()
	return b.b.String()
}
