package main

import (
	"./adapter"
)

func startMain() {
	p := adapter.NewPrintBanner("Hello")
	p.PrintWeak()
	p.PrintString()
}

func main() {
	startMain()
}
