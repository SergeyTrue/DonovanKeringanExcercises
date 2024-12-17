package main

import (
	"testing"
)

func TestUnits(t *testing.T) {
	tests := []struct {
		unit float64
		want float64
	}{
		{unit: KB, want: 1000},
		{unit: MB, want: 1000000},
		{unit: GB, want: 1000000000},
	}

	for _, test := range tests {
		if got := printUnits(test.unit); got != test.want {
			t.Errorf("printUnit(%.2f) = %.2f, want %.2f", test.unit, got, test.want)
		}
	}
}
