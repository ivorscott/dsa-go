package time

import (
	"testing"
)

func benchmarkExponential(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Exponential(i)
	}
}

func BenchmarkExponential10(b *testing.B) { benchmarkExponential(10, b) }

func BenchmarkExponential100(b *testing.B) { benchmarkExponential(100, b) }

func BenchmarkExponential1000(b *testing.B) { benchmarkExponential(1000, b) }

// func BenchmarkExponential10000(b *testing.B) {
// 	benchmarkExponential(10000, b)
// }
