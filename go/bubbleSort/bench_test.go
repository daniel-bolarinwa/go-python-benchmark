package bubbleSort

import (
	"math/rand"
	"testing"
)

func generateSlice(n int) []int {
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, rand.Intn(1e9))
	}
	return s
}

// feel free to test this with slices/arrays of varying sizes to evaluate further
func BenchmarkBubbleSort10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := generateSlice(10000)
		b.StartTimer()
		bubbleSort(s)
	}
}
