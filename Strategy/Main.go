package main

import (
	"fmt"

	"./strategy"
)

func startMain() {
	player1 := strategy.NewPlayer("Taro", strategy.NewWinningStrategy())
	player2 := strategy.NewPlayer("Hana", strategy.NewCircularStrategy())

	for i := 0; i < 10000; i++ {
		hand1 := player1.NextHand()
		hand2 := player2.NextHand()
		if hand1.IsStrongerThan(hand2) {
			fmt.Printf("Winner:%s\n", player1.ToString())
			player1.Win()
			player2.Lose()
		} else if hand1.IsWeakerThan(hand2) {
			fmt.Printf("Winner:%s\n", player2.ToString())
			player1.Lose()
			player2.Win()
		} else {
			fmt.Println("Even...")
			player1.Even()
			player2.Even()
		}
	}

	fmt.Println("Total Result:")
	fmt.Println(player1.ToString())
	fmt.Println(player2.ToString())
}

func main() {
	startMain()
}
