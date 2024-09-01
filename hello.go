package main

import (
	"fmt"
	"math"
	"runtime"
)

func Sqrt(x float64) float64 {
	z := 1.0
	tries := 100
	for i := 0; i < tries; i++ {
		previous := z
		z -= (z*z - x) / (2 * z)
		fmt.Println("try", i, ":", z)
		if closeEnough(z, previous) {
			break
		}
	}
	return z
}

func closeEnough(x, y float64) bool {
	return math.Abs(x-y) < 0.00000001
}

func main() {
// 	num := 87123.0
// 	fmt.Println("Mine: ", Sqrt(num))
	// 	fmt.Println("Go's: ", math.Sqrt(num))
	switch os:=runtime.GOOS; os {
	case "darwin":
		fmt.Println("MacOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}
}
