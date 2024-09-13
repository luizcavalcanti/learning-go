package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	ocurrences := make(map[string][]string)
	files := os.Args[1:]
	for _, filePath := range files {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(file, filePath, counts, ocurrences)
		file.Close()
	}
	printResults(counts, ocurrences)
}

func countLines(file *os.File, filename string, counts map[string]int, ocurrences map[string][]string) {
	// fmt.Println(filename)
	input := bufio.NewScanner(file)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		ocurrences[line] = append(ocurrences[line], filename)//.Append(file)
	}
	// TODO: tratar erros em input.Err()
}

func printResults(counts map[string]int, ocurrences map[string][]string) {
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
			fmt.Println(ocurrences[line])
		}
	}
}
