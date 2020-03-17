package memento

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Gamer is struct
type Gamer struct {
	fruitname, fruits []string
	money             int
}

// NewGamer func for initializing Game
func NewGamer(money int) *Gamer {
	return &Gamer{
		fruitname: []string{"リンゴ", "ぶどう", "バナナ", "みかん"},
		money:     money,
	}
}

// GetMoney func for fetching money in Gamer
func (g *Gamer) GetMoney() int {
	return g.money
}

// Bet func for betting
func (g *Gamer) Bet() {
	rand.Seed(time.Now().UnixNano())
	dice := rand.Intn(6) + 1
	if dice == 1 {
		g.money += 100
		fmt.Println("所持金が増えました")
	} else if dice == 2 {
		g.money /= 2
		fmt.Println("所持金が半分になりました")
	} else if dice == 6 {
		f := g.getFruit()
		fmt.Printf("フルーツ(%s)をもらいました\n", f)
		g.fruits = append(g.fruits, f)
	} else {
		fmt.Println("何も起こりませんでした")
	}
}

// CreateMemento func for creating Memento
func (g *Gamer) CreateMemento() *Memento {
	m := &Memento{money: g.money}
	for _, f := range g.fruits {
		if strings.HasPrefix(f, "おいしい") {
			m.addFruit(f)
		}
	}
	return m
}

// RestoreMemento func for restoring from Memento
func (g *Gamer) RestoreMemento(memento *Memento) {
	g.money = memento.money
	g.fruits = memento.GetFruits()
}

// Print func for printing current value in Gamer
func (g *Gamer) Print() string {
	return fmt.Sprintf("[money = %d, fruits = %s]", g.money, g.fruits)
}

func (g *Gamer) getFruit() string {
	prefix := ""
	if rand.Int()%2 == 0 {
		prefix = "おいしい"
	}
	return prefix + g.fruitname[rand.Intn(len(g.fruitname))]
}
