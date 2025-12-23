package main

import (
	"fmt"
	"time"
)

func counter2(max int) <-chan int {
	ch := make(chan int, 3)
	go func() {
		for i := 0; i < max; i++ {
			fmt.Println("sending: ", i)
			ch <- i
			fmt.Println("sending done: ", i)
		}
		close(ch)
	}()
	return ch
}

func main() {
	c := counter2(5)
	for v := range c {
		fmt.Println("Channel", v)
		time.Sleep(60 * time.Second)

	}
}
