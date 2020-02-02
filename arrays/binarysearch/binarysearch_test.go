package main

import "testing"

func benchmarkBinarySearch(data []int, value int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearch(data, value)
	}
}

func BenchmarkBinarySearch1(b *testing.B) { benchmarkBinarySearch([]int{1, 2, 3, 4, 5}, 5, b) }

func BenchmarkBinarySearch2(b *testing.B) {
	benchmarkBinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, b)
}

func BenchmarkBinarySearch3(b *testing.B) {
	benchmarkBinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 15, b)
}

func BenchmarkBinarySearch4(b *testing.B) {
	benchmarkBinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 20, b)
}
