package main

import (
	"./bridge"
)

func startMain() {
	d1 := bridge.DisplayFunc{Impl: &bridge.DisplayStringImpl{String: "Hello Japan"}}
	d2 := bridge.DisplayCountFunc{Impl: &bridge.DisplayStringImpl{String: "Hello Japan"}}
	d3 := bridge.DisplayCountFunc{Impl: &bridge.DisplayStringImpl{String: "Hello Universe"}}
	d4 := bridge.DisplayRandomFunc{Impl: &bridge.DisplayStringImpl{String: "Hello Universe"}}
	d5 := bridge.DisplayFunc{Impl: &bridge.DisplayTextfileImpl{FileName: "test.txt"}}
	d1.Display()
	d2.Display()
	d3.Display()
	d3.MultiDisplay(5)
	d4.RandomDisplay(5)
	d5.Display()
}

func main() {
	startMain()
}
