package main

import (
	"testing"
)

func TestMakeNameFromUrl(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"https://www.example.com/file123.pdf", "httpswwwexamplecomfile123pdf"},
		{"https://en.wikipedia.org/wiki/URL_encoding", "httpsenwikipediaorgwikiURLencoding"},
		{"https://www.google.com/search?q=golang", "httpswwwgooglecomsearchqgolang"},
		{"", ""},
		{"12345", "12345"},
		{"!@#$%^&*()", ""}}

	for _, test := range tests {
		result := MakeNameFromUrl(test.input)

		if result != test.expected {
			t.Errorf("MakeNameFromUrl(%q) = %q, want %q", test.input, result, test.expected)
		}
	}
}
