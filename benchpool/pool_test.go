package benchpool

import "testing"

func BenchmarkAppendMany(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AppendMany()
	}
}

func BenchmarkPoolAppendMany(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := PoolAppendMany()
		slicePool.Put(s[:0])
	}
}
