package main

import (
	"fmt"
	"maps"
)

func main() {
	FirstWordMap := make(map[byte]int)
	SecondWordMap := make(map[byte]int)
	s1 := "helilo"
	s2 := "elloh"
	if len(s1) != len(s2) {
		fmt.Println("Not anagrams")
		return
	}
	for i := 0; i < len(s1); i++ {
		if _, ok := FirstWordMap[s1[i]]; !ok {
			FirstWordMap[s1[i]]++
			SecondWordMap[s1[i]]++
		}
	}

	if maps.Equal(FirstWordMap, SecondWordMap) {
		fmt.Println("Anagrams")
		return
	} else {
		fmt.Println("Not anagrams")
	}

}
