package main

import "fmt"

func maxThree(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > c {
		return b
	}

	return c
}

func main() {
	fmt.Println("Max in (12, 25, 4):", maxThree(12, 25, 4))
}
