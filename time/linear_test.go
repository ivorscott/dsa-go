package time

import (
	"testing"
)

func benchmarkLinear(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Linear(i)
	}
}

func BenchmarkLinear10(b *testing.B) { benchmarkLinear(10, b) }

func BenchmarkLinear100(b *testing.B) { benchmarkLinear(100, b) }

func BenchmarkLinear1000(b *testing.B) { benchmarkLinear(1000, b) }

func BenchmarkLinear10000(b *testing.B) { benchmarkLinear(10000, b) }
