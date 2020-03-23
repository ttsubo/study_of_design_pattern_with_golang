package bridge

import (
	"fmt"
	"strings"
)

// DisplayStringImpl is struct
type DisplayStringImpl struct {
	str   string
	width int
}

// NewDisplayStringImpl func for initializing DisplayStringImpl
func NewDisplayStringImpl(str string) *DisplayStringImpl {
	return &DisplayStringImpl{
		str:   str,
		width: len(str),
	}
}

func (d *DisplayStringImpl) rawOpen() {
	d.printLine()
}

func (d *DisplayStringImpl) rawPrint() {
	fmt.Printf("|%s|\n", d.str)
}

func (d *DisplayStringImpl) rawClose() {
	d.printLine()
	fmt.Println("")
}

func (d *DisplayStringImpl) printLine() {
	line := strings.Repeat("-", d.width)
	fmt.Printf("+%s+\n", line)
}
