package main

import (
	"./adapter"
)

func startMain() {
	p := &adapter.PrintBanner{String: "Hello"}
	p.PrintWeak()
	p.PrintString()

}

func main() {
	startMain()
}
