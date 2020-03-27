package main

import (
	"flag"

	"./flyweight"
)

func startMain(str string) {
	bs := flyweight.NewBigString(str)
	bs.Print()
}

func main() {
	flag.Parse()
	startMain(flag.Arg(0))
}
