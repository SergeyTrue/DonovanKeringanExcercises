package main

import "fmt"

func main() {
	var num uint64 = 140
	fmt.Println(PopCountSingleExp(num))
	fmt.Println(PopCountLoop(num))
	fmt.Println(PopCountShift(num))
	fmt.Println(PopCountClearing(num))
}
