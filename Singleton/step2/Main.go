package main

import "fmt"

// Singleton is struct
type Singleton struct {
	input int
}

func (s *Singleton) getInstance(input int) *Singleton {
	s.input = input
	return s
}

// MyClass is struct
type MyClass struct {
	Singleton
}

var myclass = &MyClass{}

func main() {
	one := myclass.getInstance(1)
	fmt.Printf("one.input=%d\n", one.input)
	two := myclass.getInstance(2)
	fmt.Printf("one.input=%d, twe.input=%d\n", one.input, two.input)
	one.input = 0
	fmt.Printf("one.input=%d, twe.input=%d\n", one.input, two.input)
}
