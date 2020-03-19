package main

import (
	"./observer"
)

func startMain() {
	generator := observer.NewRandomNumberGenerator()
	observer1 := &observer.DigitObserver{}
	observer2 := &observer.GraphObserver{}
	generator.AddObserver(observer1)
	generator.AddObserver(observer2)
	generator.Execute()
}

func main() {
	startMain()
}
