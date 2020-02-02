package main

import "testing"

func benchmarkFactorial(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(n)
	}
}

func BenchmarkFactorial1(b *testing.B) { benchmarkFactorial(5, b) }

func BenchmarkFactorial2(b *testing.B) {
	benchmarkFactorial(10, b)
}

func BenchmarkFactorial3(b *testing.B) {
	benchmarkFactorial(15, b)
}

func BenchmarkFactorial4(b *testing.B) {
	benchmarkFactorial(20, b)
}
