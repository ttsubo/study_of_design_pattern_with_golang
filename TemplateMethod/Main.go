package main

import (
	"fmt"

	templateMethod "./templatemethod"
)

func startMain() {
	d1 := templateMethod.NewCharDisplay('H')
	d2 := templateMethod.NewStringDisplay("Hello, World!")
	d1.Display()
	fmt.Println("")
	d2.Display()
}

func main() {
	startMain()
}
