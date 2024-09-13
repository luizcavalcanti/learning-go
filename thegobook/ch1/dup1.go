// Dup1 prints the text of each line that apperas more than
// once in the standard input, preceded by its count
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}
	// TODO: tratar potenciais erros vindos de input.Err()

	fmt.Println()
	fmt.Println("Linhas repetidas:")
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
