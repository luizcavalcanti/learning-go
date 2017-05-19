package main

import (
	"fmt"
)

/*
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32, represents a Unicode code point
float32 float64
complex64 complex128
*/

var (
	Question bool = false
	Quote string = "Testing strings"
	MaxInt uint64 = 1<<64 - 1
	z complex128 = -5 + 21i
)

func print_types() {
	fmt.Println("\nRun types.go example:")
	fmt.Printf("Type: %T Value: %v\n", Question, Question)
	fmt.Printf("Type: %T Value: %v\n", Quote, Quote)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
