package main

import "testing"

func TestComma(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"1234", "1,234"},
		{"1.0", "1.0"},
		{"12.0", "12.0"},
		{"123.4356", "123.4356"},
		{"1234.4356", "1,234.4356"},
	}
	for _, test := range tests {
		if got := comma(test.input); got != test.expected {
			t.Errorf("Input = %s , expected =  %s, got = %s", test.input, test.expected, got)
		}
	}
}
