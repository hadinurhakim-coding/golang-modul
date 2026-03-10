package mathutil

import "testing"

func BenchmarkSum_1000(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Sum(nums)
	}
}

