package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
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
func main() {
	fmt.Println(comma("123456578"))
}
