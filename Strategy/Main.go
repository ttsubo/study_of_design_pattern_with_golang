package main

import (
	"fmt"

	"./strategy"
)

func main() {
	player1 := strategy.Player{Name: "Taro", PlayerStrategy: &strategy.WinningStrategy{}}
	player2 := strategy.Player{Name: "Hana", PlayerStrategy: &strategy.CircularStrategy{}}

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
	fmt.Printf("%s\n", player1.ToString())
	fmt.Printf("%s\n", player2.ToString())
}
