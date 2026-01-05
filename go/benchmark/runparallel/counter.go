package runparallel

import (
	"sync"
	"sync/atomic"
)

type Counter struct {
	v int64
}

func (c *Counter) Add(n int64) {
	c.v += n
}

type MutexCounter struct {
	v  int64
	mu sync.Mutex
}

func (c *MutexCounter) Add(n int64) {
	c.mu.Lock()
	c.v += n
	c.mu.Unlock()
}

type AtomicCounter struct {
	v int64
}

func (c *AtomicCounter) Add(n int64) {
	atomic.AddInt64(&c.v, n)
}
