package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode"
)

func MakeNameFromUrl(url string) string {
	var sb strings.Builder
	for _, char := range url {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			sb.WriteRune(char)
		}
	}
	return sb.String()
}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	f, err := os.Create(MakeNameFromUrl(url))
	if err != nil {
		ch <- fmt.Sprint("Error creating file: ", err)
	}
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint("Error fetching " + url)
	}
	defer resp.Body.Close()
	nbytes, err := io.Copy(f, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching " + url)
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f\t%s\t%10d", secs, url, nbytes)

}

func main() {
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)

	}
	fmt.Println(<-ch)
}
