package performance

import "testing"

func BenchmarkMyFunc(b *testing.B) {
	ints := []int{}
	for i := 0; i < 1000; i++ {
		ints = append(ints, i)
	}

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		MyFunc(ints)
	}
}
