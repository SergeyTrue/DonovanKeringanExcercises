package main

import "testing"

func TestRotateSinglePass(t *testing.T) {
	tests := []struct {
		input []int
		k     int
		want  []int
	}{
		{
			[]int{1, 2, 3, 4, 5}, 4, []int{4, 5, 1, 2, 3},
		},
	}

}
