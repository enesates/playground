package benchmark

import (
	"math/rand"
	"slices"
)

func QuickSort(n []int) {
	slices.Sort(n)
}

func BubbleSort(n []int) {
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n)-i-1; j++ {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			}
		}
	}
}
func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(1000000)
	}
	return slice
}
