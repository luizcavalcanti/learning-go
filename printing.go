package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
)

func print_printing() {
	// Setting "random" seed to rand package
	rand.Seed(time.Now().Unix())

	fmt.Println("\nRun printing.go example:")
	fmt.Printf("Hello, World\n")
	fmt.Printf("The current time is %s\n", time.Now().String())
	fmt.Printf("A (not so) random number: %d\n", rand.Intn(10))
	fmt.Printf("And pi, of course: %f\n", math.Pi)
	fmt.Printf("To print a empty string, use %%q: %q\n", "")
}
