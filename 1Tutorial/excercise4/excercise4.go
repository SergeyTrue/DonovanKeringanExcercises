package main

import (
	"fmt"
	"os"
	"strings"
)

type RepeatedLineInfo struct {
	fileName  string
	line      string
	lineCount int
}

func (lineInfo *RepeatedLineInfo) PrintLineInfo() {
	fmt.Printf("%s\t%s\t%d\n",
		lineInfo.fileName,
		lineInfo.line,
		lineInfo.lineCount)
}

func main() {
	fileStats := []RepeatedLineInfo{}

	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)

		if err != nil {
			fmt.Println(err)
		}

		lines := strings.Split(string(data), "\r\n")
		counter := make(map[string]int)
		for _, line := range lines {
			counter[line]++
		}

		for line, count := range counter {
			if count > 1 {
				fileLineInfos := RepeatedLineInfo{
					fileName:  filename,
					line:      line,
					lineCount: count,
				}
				fileStats = append(fileStats, fileLineInfos)
			}

		}
	}

	for _, entry := range fileStats {
		entry.PrintLineInfo()
	}
}
