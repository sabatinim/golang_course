package bench_test_example

import (
	"average_pkg"
	"testing"
)

func BenchmarkTestAverage(b *testing.B) {
	b.N = 10000
	for i := 0; i < b.N; i++ {
		var v float64
		v = average_pkg.Average([]float64{1, 2})
		if v != 1.5 {
			b.Error("Expected 1.5, got ", v)
		}
	}
}
