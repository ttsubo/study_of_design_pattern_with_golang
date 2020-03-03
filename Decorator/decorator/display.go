package decorator

import "fmt"

// Display is struct
type Display struct {
}

// DisplayInterface is interface
type DisplayInterface interface {
	getColumns() int
	getRows() int
	getRowText(row int) string
}

// Show func for printing something
func (d *Display) Show(display DisplayInterface) {
	str := ""
	for i := 0; i < display.getRows(); i++ {
		str += display.getRowText(i) + "\n"
	}
	fmt.Printf("%s", str)
}

// StringDisplay is struct
type StringDisplay struct {
	*Display
	String string
}

func (s *StringDisplay) getColumns() int {
	return len(s.String)
}

func (s *StringDisplay) getRows() int {
	return 1
}

func (s *StringDisplay) getRowText(row int) string {
	if row == 0 {
		return s.String
	}
	return ""
}

// NewStringDisplay func for initalizing StringDisplay
func NewStringDisplay(str string) *StringDisplay {
	return &StringDisplay{
		Display: &Display{},
		String:  str,
	}
}
