package main

import "fmt"

func reverse(arr *[5]int, start, end int) {
	if start < 0 || end >= len(arr) || start >= end {
		return
	}
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
func main() {

	a := [5]int{1, 2, 3, 4, 5}

	fmt.Println(a)
	reverse(&a, 1, 3)

	fmt.Println(a)
}
