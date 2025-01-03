package main

import (
	"crypto/sha256"
	"testing"
)

func TestShaBitDiff(t *testing.T) {
	tests := []struct {
		c1   [32]byte
		c2   [32]byte
		want int
	}{
		{sha256.Sum256([]byte("x")), sha256.Sum256([]byte("X")), 122},
	}

	for _, tt := range tests {
		got := ShaBitDiff(tt.c1, tt.c2)
		if got != tt.want {
			t.Errorf("ShaBitDiff(%v,%v) = %d, want %d", tt.c1, tt.c2, got, tt.want)
		}
	}
}
