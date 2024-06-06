package excercise3

import "testing"

func BenchmarkJoinArgs(b *testing.B) {
	args := []string{"b", "a", "c"}
	for i := 0; i < b.N; i++ {
		JoinArgs(args)
	}
}

func BenchmarkConcatArgs(b *testing.B) {
	args := []string{"b", "a", "c"}
	for i := 0; i < b.N; i++ {
		ConcatArgs(args)
	}
}
