/*
Exercise 4.1: Write a function that counts the number of bits that are different in two SHA256
hashes. (See PopCount from Section 2.6.2.)
*/

package main

import (
	"crypto/sha256"
	"fmt"
)

func ShaBitDiff(c1 [32]byte, c2 [32]byte) int {
	var count int
	var resArr [32]byte
	for i := range len(c1) {
		resArr[i] = c2[i] ^ c1[i]
	}

	for i := 0; i < len(resArr); i++ {
		for j := 0; j < 8; j++ {
			count += int(resArr[i] >> (j) & 1)
		}

	}
	return count
}

func main() {

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(ShaBitDiff(c1, c2))

}
