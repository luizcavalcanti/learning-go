// Echo3 prints its command-line arguments
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	withForLoop()
	withJoin()
}

func withForLoop() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println("With for loop:", time.Since(start).Seconds())

}
func withJoin() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("With strings.Join:", time.Since(start).Seconds())	
}
