package strategy

// Const Value for HandGame
const (
	HandValueGUU = iota
	HandValueCHO
	HandValuePAA
)

//Hand is struct
type Hand struct {
	handValue int
}

var hands []*Hand

func init() {
	hands = []*Hand{
		&Hand{HandValueGUU},
		&Hand{HandValueCHO},
		&Hand{HandValuePAA},
	}
}

func getHand(handValue int) *Hand {
	return hands[handValue]
}

// IsStrongerThan func can judge result of handGame as Winner
func (myHand *Hand) IsStrongerThan(opponentHand *Hand) bool {
	return myHand.fight(opponentHand) == 1
}

// IsWeakerThan func can judge result of handGame as Looser
func (myHand *Hand) IsWeakerThan(opponentHand *Hand) bool {
	return myHand.fight(opponentHand) == -1
}

func (myHand *Hand) fight(opponentHand *Hand) int {
	if myHand == opponentHand {
		return 0
	} else if (myHand.handValue+1)%3 == opponentHand.handValue {
		return 1
	} else {
		return -1
	}
}
