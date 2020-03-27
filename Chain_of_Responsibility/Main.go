package main

import (
	chainOfResponsibility "./chain_of_responsibility"
)

func startMain() {
	alice := chainOfResponsibility.NewNoSupport("Alice")
	bob := chainOfResponsibility.NewLimitSupport("Bob", 100)
	charlie := chainOfResponsibility.NewSpecialSupport("Charlie", 429)
	diana := chainOfResponsibility.NewLimitSupport("Diana", 200)
	elmo := chainOfResponsibility.NewOddSupport("Elmo")
	fred := chainOfResponsibility.NewLimitSupport("Fred", 300)

	alice.SetNext(bob).SetNext(charlie).SetNext(diana).SetNext(elmo).SetNext(fred)

	for i := 0; i < 500; i += 33 {
		alice.Handle(chainOfResponsibility.NewTrouble(i))
	}
}

func main() {
	startMain()
}
