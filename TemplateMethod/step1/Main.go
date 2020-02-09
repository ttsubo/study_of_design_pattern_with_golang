package main

import (
	"fmt"

	templateMethod "./templatemethod"
)

func main() {
	d1 := &templateMethod.CharDisplay{Char: 'H'}
	d2 := &templateMethod.StringDisplay{Str: "Hello, World!"}
	templateMethod.Display(d1)
	fmt.Println("")
	templateMethod.Display(d2)
}
