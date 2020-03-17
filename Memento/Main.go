package main

import (
	"fmt"
	"time"

	"./memento"
)

func startMain() {
	gamer := memento.NewGamer(100)
	memento := gamer.CreateMemento()

	for i := 0; i < 100; i++ {
		fmt.Printf("==== %d\n", i)
		fmt.Printf("現状:%s\n", gamer.Print())
		gamer.Bet()
		fmt.Printf("所持金は%d円になりました\n", gamer.GetMoney())

		if gamer.GetMoney() > memento.GetMoney() {
			fmt.Println("      (だいぶ増えたので、現在の状態を保存しておこう)")
			memento = gamer.CreateMemento()
		} else if gamer.GetMoney() < memento.GetMoney()/2 {
			fmt.Println("      (だいぶ減ったので、以前の状態に復帰しよう)")
			gamer.RestoreMemento(memento)
		}

		time.Sleep(time.Second * 1)
		fmt.Println("")
	}
}

func main() {
	startMain()
}
