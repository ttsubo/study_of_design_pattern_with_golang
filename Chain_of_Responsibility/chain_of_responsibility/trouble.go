package chainOfResponsibility

import "fmt"

// Trouble is struct
type Trouble struct {
	number int
}

// NewTrouble func for initializing Trouble
func NewTrouble(number int) *Trouble {
	return &Trouble{
		number: number,
	}
}

func (t *Trouble) getNumber() int {
	return t.number
}

func (t *Trouble) print() string {
	return fmt.Sprintf("[Trouble %d]", t.number)
}
