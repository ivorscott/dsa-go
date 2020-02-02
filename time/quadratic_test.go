package time

import (
	"testing"
)

func benchmarkQuadratic(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Quadratic(i)
	}
}

func BenchmarkQuadratic10(b *testing.B) { benchmarkQuadratic(10, b) }

func BenchmarkQuadratic100(b *testing.B) { benchmarkQuadratic(100, b) }

func BenchmarkQuadratic1000(b *testing.B) { benchmarkQuadratic(1000, b) }

func BenchmarkQuadratic10000(b *testing.B) {
	benchmarkQuadratic(10000, b)
}
