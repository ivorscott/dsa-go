package main

import "testing"

func benchmarkFindOne(data []int, value int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindOne(data, value)
	}
}

func BenchmarkFindOne1(b *testing.B) { benchmarkFindOne([]int{1, 2, 3, 4, 5}, 5, b) }

func BenchmarkFindOne2(b *testing.B) { benchmarkFindOne([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, b) }

func BenchmarkFindOne3(b *testing.B) {
	benchmarkFindOne([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 15, b)
}

func BenchmarkFindOne4(b *testing.B) {
	benchmarkFindOne([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 20, b)
}
