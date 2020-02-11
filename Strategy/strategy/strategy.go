package strategy

import (
	"math/rand"
	"time"
)

// Strategy is interface
type Strategy interface {
	NextHand() *Hand
	study(win bool)
}

// WinningStrategy is struct
type WinningStrategy struct {
	won      bool
	prevHand *Hand
}

// NextHand func can handle result of preHand
func (ws *WinningStrategy) NextHand() *Hand {
	if !ws.won {
		rand.Seed(time.Now().UnixNano())
		ws.prevHand = getHand(rand.Intn(3))
	}
	return ws.prevHand
}

func (ws *WinningStrategy) study(win bool) {
	ws.won = win
}

// CircularStrategy is struct
type CircularStrategy struct {
	hand int
}

// NextHand func can handle result of each Hand
func (cs *CircularStrategy) NextHand() *Hand {
	return getHand(cs.hand)
}

func (cs *CircularStrategy) study(win bool) {
	cs.hand = (cs.hand + 1) % 3
}
