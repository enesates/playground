package main

import "fmt"

func main() {
	original := [5]int{1, 2, 3, 4, 5}

	window := original[1:3]
	fmt.Println(window)

	window[0] = 999
	fmt.Println(window)
	fmt.Println(original)

	///////////////////////////

	original2 := []int{1, 2, 3, 4, 5}
	window2 := make([]int, len(original2))

	copy(window2, original2)
	window2[0] = 999
	fmt.Println(original2)
}
