package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error fetching %s\n", url)
			continue
		}
		defer resp.Body.Close()
		body, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Printf("Error reading %s\n", url)
		}
		fmt.Println(body)
	}
}
