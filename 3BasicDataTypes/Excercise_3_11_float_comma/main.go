package main

import (
	"bytes"
	"fmt"
	"strings"
)

func split(s string) (string, string) {
	if dot := strings.LastIndex(s, "."); dot != -1 {
		integral := s[:dot]
		decimal := s[dot:]
		return integral, decimal
	}
	return s, ""
}

func int_comma(s string) string {

	n := len(s)
	if n <= 3 {
		return s
	}
	buf := bytes.Buffer{}
	pre := n % 3
	if pre > 0 {
		buf.WriteString(s[:pre])
		buf.WriteString(",")
	}
	for i := pre; i < n-pre; i += 3 {
		buf.WriteString(s[i : i+3])
		if n-i > 3 {
			buf.WriteString(",")
		}
	}
	return buf.String()
}

func comma(s string) string {
	integer, decimal := split(s)
	return int_comma(integer) + decimal
}

func main() {
	fmt.Println(comma("123456789"))
}
