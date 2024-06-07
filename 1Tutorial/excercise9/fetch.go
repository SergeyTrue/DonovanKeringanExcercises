package main

import (
	"fmt"
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
			fmt.Printf("Error fetching url %v\n", err)
		}
		defer resp.Body.Close()
		status := resp.Status
		fmt.Printf("Status %v\n", status)
	}
}
