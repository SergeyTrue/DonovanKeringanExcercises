package _Tutorial

import (
	"fmt"
	"os"
)

func main() {

	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i, ":", os.Args[i])
	}
}
