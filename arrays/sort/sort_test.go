package main

import "testing"

func benchmarkSortA(d []int, size int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortA(d, size)
	}
}

func benchmarkSortB(d []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortB(d)
	}
}

func BenchmarkSortA1(b *testing.B) { benchmarkSortA([]int{2, 1, 3, 4, 6, 5}, 5, b) }

func BenchmarkSortA2(b *testing.B) {
	benchmarkSortA([]int{9, 7, 1, 4, 2, 3, 8, 5, 6, 10}, 10, b)
}

func BenchmarkSortA3(b *testing.B) {
	benchmarkSortA([]int{11, 12, 1, 4, 2, 14, 3, 8, 15, 5, 6, 7, 13, 9, 10}, 15, b)
}

func BenchmarkSortA4(b *testing.B) {
	benchmarkSortA([]int{2, 1, 19, 3, 4, 8, 6, 5, 7, 9, 11, 12, 13, 15, 16, 10, 17, 18, 14, 20}, 20, b)
}

// Sort 2

func BenchmarkSortB1(b *testing.B) { benchmarkSortB([]int{2, 1, 3, 4, 6, 5}, b) }

func BenchmarkSortB2(b *testing.B) {
	benchmarkSortB([]int{9, 7, 1, 4, 2, 3, 8, 5, 6, 10}, b)
}

func BenchmarkSortB3(b *testing.B) {
	benchmarkSortB([]int{11, 12, 1, 4, 2, 14, 3, 8, 15, 5, 6, 7, 13, 9, 10}, b)
}

func BenchmarkSortB4(b *testing.B) {
	benchmarkSortB([]int{2, 1, 19, 3, 4, 8, 6, 5, 7, 9, 11, 12, 13, 15, 16, 10, 17, 18, 14, 20}, b)
}
