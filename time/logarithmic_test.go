package time

import (
	"testing"
)

func benchmarkLogarithmic(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Logarithmic(i)
	}
}

func BenchmarkLogarithmic10(b *testing.B) { benchmarkLogarithmic(10, b) }

func BenchmarkLogarithmic100(b *testing.B) { benchmarkLogarithmic(100, b) }

func BenchmarkLogarithmic1000(b *testing.B) { benchmarkLogarithmic(1000, b) }

func BenchmarkLogarithmic10000(b *testing.B) { benchmarkLogarithmic(10000, b) }
