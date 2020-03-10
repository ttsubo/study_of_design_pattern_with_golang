package chainOfResponsibility

import "fmt"

// Trouble is struct
type Trouble struct {
	Number int
}

func (t *Trouble) getNumber() int {
	return t.Number
}

func (t *Trouble) print() string {
	return fmt.Sprintf("[Trouble %d]", t.Number)
}
