package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Fetch Error: %s\n", err)
		}
		defer resp.Body.Close()
		body, err := io.Copy(os.Stdout, resp.Body)
		fmt.Println(body)
	}
}
