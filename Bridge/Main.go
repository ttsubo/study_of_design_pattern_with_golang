package main

import (
	"./bridge"
)

func startMain() {
	d1 := bridge.NewDisplayFunc(bridge.NewDisplayStringImpl("Hello Japan"))
	d2 := bridge.NewDisplayCountFunc(bridge.NewDisplayStringImpl("Hello Japan"))
	d3 := bridge.NewDisplayCountFunc(bridge.NewDisplayStringImpl("Hello Universe"))
	d4 := bridge.NewDisplayRandomFunc(bridge.NewDisplayStringImpl("Hello Universe"))
	d5 := bridge.NewDisplayFunc(bridge.NewDisplayTextfileImpl("test.txt"))
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
