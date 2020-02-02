package main

import "testing"

func benchmarkGCD(m int, n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCD(m, n)
	}
}

func BenchmarkGCD1(b *testing.B) { benchmarkGCD(24, 18, b) }

func BenchmarkGCD2(b *testing.B) {
	benchmarkGCD(1550, 200, b)
}

func BenchmarkGCD3(b *testing.B) {
	benchmarkGCD(12400, 224, b)
}

func BenchmarkGCD4(b *testing.B) {
	benchmarkGCD(440224, 224, b)
}
