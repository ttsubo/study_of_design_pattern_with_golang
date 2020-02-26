package main

import "fmt"

// Singleton is struct
type Singleton struct {
}

// MyClass is struct
type MyClass struct {
	Singleton
	input int
}

var instance *MyClass

func getInstance(input int) *MyClass {
	if instance == nil {
		instance = &MyClass{input: input}
	} else {
		instance.input = input
	}
	return instance
}

func main() {
	one := getInstance(1)
	fmt.Printf("one.input=%d\n", one.input)
	two := getInstance(2)
	fmt.Printf("one.input=%d, twe.input=%d\n", one.input, two.input)
	one.input = 0
	fmt.Printf("one.input=%d, twe.input=%d\n", one.input, two.input)
}
