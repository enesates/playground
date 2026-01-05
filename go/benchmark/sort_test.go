package benchmark

import (
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {
	data := generateRandomSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort(data)
	}
}
func BenchmarkBubbleSortSmall(b *testing.B) {
	data := generateRandomSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort(data)
	}
}
func BenchmarkBubbleSortLarge(b *testing.B) {
	data := generateRandomSlice(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort(data)
	}
}
func BenchmarkQuickSort1(b *testing.B) {
	data := generateRandomSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(data)
	}
}
func BenchmarkQuickSortSmall(b *testing.B) {
	data := generateRandomSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(data)
	}
}
func BenchmarkQuickSortLarge(b *testing.B) {
	data := generateRandomSlice(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(data)
	}
}
