package main

import (
	"./observer"
)

func startMain() {
	generator := observer.NewRandomNumberGenerator()
	observer1 := observer.NewDigitObserver()
	observer2 := observer.NewGraphObserver()
	generator.AddObserver(observer1)
	generator.AddObserver(observer2)
	generator.Execute()
}

func main() {
	startMain()
}
