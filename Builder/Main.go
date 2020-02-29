package main

import (
	"flag"
	"fmt"

	"./builder"
)

func startMain(opt string) {
	if opt == "plain" {
		builderObj := builder.TextBuilder{}
		director := builder.Director{Builder: &builderObj}
		director.Construct()
		result := builderObj.GetResult()
		fmt.Println(result)
	} else if opt == "html" {
		builderObj := builder.HTMLBuilder{}
		director := builder.Director{Builder: &builderObj}
		director.Construct()
		result := builderObj.GetResult()
		fmt.Printf("[%s] was created.\n", result)
	}
}

func main() {
	flag.Parse()
	startMain(flag.Arg(0))
}
