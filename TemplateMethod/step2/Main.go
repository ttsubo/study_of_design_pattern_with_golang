package main

import (
	"fmt"

	templateMethod "./templatemethod"
)

func main() {
	d1 := &templateMethod.CharDisplay{Char: 'H'}
	d2 := &templateMethod.StringDisplay{Str: "Hello, World!"}
	d1.Display(d1)
	fmt.Println("")
	d2.Display(d2)
}
