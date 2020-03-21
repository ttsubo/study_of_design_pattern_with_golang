package templateMethod

import "fmt"

// Printer is interface
type Printer interface {
	open()
	print()
	close()
}

// AbstractDisplay is struct
type AbstractDisplay struct {
	printer Printer
}

// Display for printout result
func (a *AbstractDisplay) Display() {
	a.printer.open()
	for i := 0; i < 5; i++ {
		a.printer.print()
	}
	a.printer.close()
}

// CharDisplay is struct
type CharDisplay struct {
	*AbstractDisplay
	ch byte
}

// NewCharDisplay func for initializing CharDisplay
func NewCharDisplay(ch byte) *CharDisplay {
	charDisplay := &CharDisplay{
		AbstractDisplay: &AbstractDisplay{},
		ch:              ch,
	}
	charDisplay.printer = charDisplay
	return charDisplay
}

func (c *CharDisplay) open() {
	fmt.Print("<<")
}
func (c *CharDisplay) print() {
	fmt.Print(string(c.ch))
}
func (c *CharDisplay) close() {
	fmt.Println(">>")
}

// StringDisplay is struct
type StringDisplay struct {
	*AbstractDisplay
	str   string
	width int
}

// NewStringDisplay func for initalizing StringDisplay
func NewStringDisplay(str string) *StringDisplay {
	stringDisplay := &StringDisplay{
		AbstractDisplay: &AbstractDisplay{},
		str:             str,
		width:           len(str),
	}
	stringDisplay.printer = stringDisplay
	return stringDisplay
}

func (s *StringDisplay) open() {
	s.printLine()
}
func (s *StringDisplay) print() {
	fmt.Printf("%s%s%s\n", "|", s.str, "|")
}
func (s *StringDisplay) close() {
	s.printLine()
}

func (s *StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < s.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
