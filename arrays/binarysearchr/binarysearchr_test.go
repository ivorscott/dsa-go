package main

import "testing"

func benchmarkBinarySearchRecursive(data []int, value int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearchRecursive(data, value)
	}
}

func BenchmarkBinarySearchRecursive1(b *testing.B) {
	benchmarkBinarySearchRecursive([]int{1, 2, 3, 4, 5}, 5, b)
}

func BenchmarkBinarySearchRecursive2(b *testing.B) {
	benchmarkBinarySearchRecursive([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, b)
}

func BenchmarkBinarySearchRecursive3(b *testing.B) {
	benchmarkBinarySearchRecursive([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 15, b)
}

func BenchmarkBinarySearchRecursive4(b *testing.B) {
	benchmarkBinarySearchRecursive([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 20, b)
}
