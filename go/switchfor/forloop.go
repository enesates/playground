package main

import "fmt"

func searchInSlice(val int, list []int) int {
	for i, elem := range list {
		if elem == val {
			return i
		}
	}

	return -1
}

func main() {
	val := 3
	list := []int{1, 2, 3}

	index := searchInSlice(val, list)

	if index >= 0 {
		fmt.Printf("%v in the position %v", val, index)
	} else {
		fmt.Println(val, "not found")
	}
}
