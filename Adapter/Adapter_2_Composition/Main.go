package main

import (
	"./adapter"
)

func startMain() {
	p := &adapter.PrintBanner{String: "Hello"}
	p.PrintWeak(p)
	p.PrintString(p)
}

func main() {
	startMain()
}
