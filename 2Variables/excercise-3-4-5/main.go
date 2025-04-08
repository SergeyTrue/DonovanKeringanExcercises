package main

import "fmt"

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	var num uint64 = 140
	for i := range pc {
		fmt.Printf("decimal i=%d, binary i=%b, pc[i]=%d, byte(i&1)=%b\n", i, i, pc[i], byte(i&1))
	}

	fmt.Println(PopCountSingleExp(num))
	fmt.Println(PopCountLoop(num))
	fmt.Println(PopCountShift(num))
	fmt.Println(PopCountClearing(num))
}
