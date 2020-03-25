package decorator

import "fmt"

type displayInterface interface {
	getColumns() int
	getRows() int
	getRowText(row int) string
}

type display struct {
	myDisplay displayInterface
}

// Show func for printing something
func (d *display) Show() {
	str := ""
	for i := 0; i < d.myDisplay.getRows(); i++ {
		str += d.myDisplay.getRowText(i) + "\n"
	}
	fmt.Printf("%s", str)
}

// StringDisplay is struct
type StringDisplay struct {
	*display
	String string
}

// NewStringDisplay func for initalizing StringDisplay
func NewStringDisplay(str string) *StringDisplay {
	stringDisplay := &StringDisplay{
		display: &display{},
		String:  str,
	}
	stringDisplay.myDisplay = stringDisplay
	return stringDisplay
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
