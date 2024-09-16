// Ftoc prints two fahrenheit-to-celsius conversions.
package main

import "fmt"

func main() {
	const freezeingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %gC\n", freezeingF, fToC(freezeingF))
	fmt.Printf("%gF = %gC\n", boilingF, fToC(boilingF))

	x := 1
	p := &x
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(&x)
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
