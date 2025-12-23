package main

import "fmt"

func countDown() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {
	c := countDown()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
}
