package main

import "testing"

func BenchmarkPopCountSingleExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountSingleExp(0xdeadbeef)
	}
}
func BenchmarkPopCountClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClearing(0xdeadbeef)
	}
}
func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(0xdeadbeef)
	}
}
