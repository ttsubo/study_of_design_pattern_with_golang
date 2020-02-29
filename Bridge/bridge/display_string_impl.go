package bridge

import (
	"fmt"
	"strings"
)

// DisplayStringImpl is struct
type DisplayStringImpl struct {
	String string
}

func (d *DisplayStringImpl) rawOpen() {
	d.printLine()
}

func (d *DisplayStringImpl) rawPrint() {
	fmt.Printf("|%s|\n", d.String)
}

func (d *DisplayStringImpl) rawClose() {
	d.printLine()
	fmt.Println("")
}

func (d *DisplayStringImpl) printLine() {
	line := strings.Repeat("-", len(d.String))
	fmt.Printf("+%s+\n", line)
}
