package main

import "testing"

func benchmarkSumArray(d []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumArray(d)
	}
}

func BenchmarkSumArray1(b *testing.B) { benchmarkSumArray([]int{1, 2, 3, 4, 5}, b) }

func BenchmarkSumArray2(b *testing.B) { benchmarkSumArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, b) }

func BenchmarkSumArray3(b *testing.B) {
	benchmarkSumArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, b)
}

func BenchmarkSumArray4(b *testing.B) {
	benchmarkSumArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, b)
}
