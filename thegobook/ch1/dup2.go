// Dup2 prints the count and text of lines that appera more than once
// in the input. It reads from stin of from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, filePath := range files {
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(file, counts)
			file.Close()
		}
	}
	printResults(counts)
}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
	// TODO: tratar erros em input.Err()
}

func printResults(counts map[string]int) {
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
