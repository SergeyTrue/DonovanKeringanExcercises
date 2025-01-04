package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

var sha = flag.String("sha", "256", "hash algorithm")

func computeHash(algorithm string, input []byte) ([]byte, error) {
	switch *sha {
	case "256":
		hash := sha256.Sum256(input)
		return hash[:], nil
	case "512":
		hash := sha512.Sum512(input)
		return hash[:], nil
	case "384":
		hash := sha512.Sum384(input)
		return hash[:], nil
	default:
		return nil, fmt.Errorf("Unknown hash algorithm: %s, Use 256, 384 or 512 instead\n", algorithm)

	}
}
func main() {
	flag.Parse()
	var input []byte
	var err error

	if len(flag.Args()) == 0 {
		input, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
			os.Exit(1)
		}
	} else {
		input = []byte(flag.Arg(0))
	}

	result, err := computeHash(*sha, input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error computing hash: %v\n", err)
	} else {
		fmt.Printf("%x\n", result)
	}
}
