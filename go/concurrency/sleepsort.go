package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var list = []int{5, 2, 10, 3, 1, 4, 6, 9, 0, 8, 7}

	wg := sync.WaitGroup{}

	for i := 0; i < len(list); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(i) * 100 * time.Millisecond)
			fmt.Printf("i am %d\n", i)
		}()
	}

	wg.Wait()
}
