package templateMethod

import "fmt"

// Printer is interface
type Printer interface {
	open()
	print()
	close()
}

// Display for printout result
func Display(p Printer) {
	p.open()
	for i := 0; i < 5; i++ {
		p.print()
	}
	p.close()
}

// CharDisplay is struct
type CharDisplay struct {
	Char byte
}

func (c *CharDisplay) open() {
	fmt.Print("<<")
}
func (c *CharDisplay) print() {
	fmt.Print(string(c.Char))
}
func (c *CharDisplay) close() {
	fmt.Println(">>")
}

// StringDisplay is struct
type StringDisplay struct {
	Str string
}

func (s *StringDisplay) open() {
	s.printLine()
}
func (s *StringDisplay) print() {
	fmt.Printf("%s%s%s\n", "|", s.Str, "|")
}
func (s *StringDisplay) close() {
	s.printLine()
}

func (s *StringDisplay) printLine() {
	fmt.Print("+")
	for _ = range s.Str {
		fmt.Print("-")
	}
	fmt.Println("+")
}
