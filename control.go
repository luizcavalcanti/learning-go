package main

import (
	"fmt"
	"runtime"
)

func testFor() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)	
	}
}

func testWhile() {
	var i int
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func testIf(test bool) {
	if test {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func testSwitch() string {
	os := runtime.GOOS
	switch os {
	case "darwin":
		return "MacOS"
	case "linux":
		return "Linux"
	default:
		return os
	}
}

func testDefer() {
	fmt.Println("proving 'defer' is stored in stack")
	for i:=0; i<5; i++ {
		defer fmt.Printf("deferred stack: %d\n", i)
	}
	fmt.Println("last line in function")
}


func printControls() {
	fmt.Println("\nRun control.do examples:")

	fmt.Println("For:")
	testFor()
	
	fmt.Println("\"While\":")
	testWhile()

	fmt.Println("If:")
	testIf(true)
	testIf(false)

	fmt.Printf("Switch: GOOS=%s\n", testSwitch())

	testDefer()
}
