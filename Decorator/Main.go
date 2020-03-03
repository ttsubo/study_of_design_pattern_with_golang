package main

import (
	"fmt"

	"./decorator"
)

func startMain() {
	b1 := decorator.NewStringDisplay("Hello, world.")
	b2 := decorator.NewSideBorder(b1, "#")
	b3 := decorator.NewFullBorder(b2)
	b1.Show(b1)
	b2.Show(b2)
	b3.Show(b3)
	fmt.Println("")
	b4 := decorator.NewSideBorder(
		decorator.NewFullBorder(
			decorator.NewFullBorder(
				decorator.NewSideBorder(
					decorator.NewFullBorder(
						decorator.NewStringDisplay("HELLO"),
					),
					"*",
				),
			),
		),
		"/",
	)
	b4.Show(b4)
}

func main() {
	startMain()
}
