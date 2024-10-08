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
		{"12345", "12,345"},
		{"123456", "123,456"},
		{"1234567", "1,234,567"},
		{"12345678", "12,345,678"},
		{"123456789", "123,456,789"},
	}
	for _, test := range tests {
		got := comma(test.input)
		if got != test.expected {
			t.Errorf("comma(%q) = %q, want %q", test.input, got, test.expected)
		}
	}
}
