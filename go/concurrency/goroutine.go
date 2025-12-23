package main

import (
	"fmt"
	"sync"
)

func MyIndex(i int) {
	fmt.Printf("i am %d\n", i)
}

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			MyIndex(i)
			wg.Done()
		}()
	}

	wg.Wait()
}
