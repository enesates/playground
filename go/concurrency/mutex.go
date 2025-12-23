package main

import "sync"

func main() {}

var (
	counter int
	mu      sync.Mutex
)

func inc() {
	mu.Lock()
	counter++
	mu.Unlock()
}
