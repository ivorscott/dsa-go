package main

import "testing"

func benchmarkFibonacci(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(n)
	}
}

func BenchmarkFibonacci1(b *testing.B) { benchmarkFibonacci(5, b) }

func BenchmarkFibonacci2(b *testing.B) {
	benchmarkFibonacci(10, b)
}

func BenchmarkFibonacci3(b *testing.B) {
	benchmarkFibonacci(15, b)
}

func BenchmarkFibonacci4(b *testing.B) {
	benchmarkFibonacci(20, b)
}
