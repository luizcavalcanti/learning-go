package main

import "fmt"

var global_var = "yes"

// Sequential parameters of the same type only
// need one type declaration
func add(x, y int) int {
	var sum int
	sum = x + y
	return sum
}

// Function with multiple return values
func swap(a, b string) (string, string) {
	return b, a
}

// Function with named return values,
// no need for explicit return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func print_functions() {
	fmt.Println("\nRun functions.go example:")

	fmt.Printf("The sum between %d and %d is %d\n", 42, 13, add(42,13))

	a, b := swap("first","second")
	fmt.Printf("Swapped strings: %s, %s\n", a, b)

	fmt.Println(split(40))

	var local_var int
	fmt.Println(local_var, global_var)
}
