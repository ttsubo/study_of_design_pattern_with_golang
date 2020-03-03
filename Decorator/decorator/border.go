package decorator

// Border is struct
type Border struct {
	*Display
	display DisplayInterface
}

// SideBorder is struct
type SideBorder struct {
	*Border
	BorderChar string
}

func (s *SideBorder) getColumns() int {
	return len(s.BorderChar)*2 + s.display.getColumns()
}

func (s *SideBorder) getRows() int {
	return s.display.getRows()
}

func (s *SideBorder) getRowText(row int) string {
	return s.BorderChar + s.display.getRowText(row) + s.BorderChar
}

// NewSideBorder func for initalizing SideBorder
func NewSideBorder(display DisplayInterface, borderChar string) *SideBorder {
	return &SideBorder{
		Border:     &Border{display: display},
		BorderChar: borderChar,
	}
}

// FullBorder is struct
type FullBorder struct {
	*SideBorder
}

func (f *FullBorder) getColumns() int {
	return 2 + f.display.getColumns()
}

func (f *FullBorder) getRows() int {
	return 2 + f.display.getRows()
}

func (f *FullBorder) getRowText(row int) string {
	if row == 0 {
		return "+" + f.makeLine("-", f.display.getColumns()) + "+"
	} else if row == f.display.getRows()+1 {
		return "+" + f.makeLine("-", f.display.getColumns()) + "+"
	} else {
		return "|" + f.display.getRowText(row-1) + "|"
	}
}

func (f *FullBorder) makeLine(char string, count int) string {
	buf := ""
	for i := 0; i < count; i++ {
		buf += char
	}
	return buf

}

// NewFullBorder func for initalizing SideBorder
func NewFullBorder(display DisplayInterface) *FullBorder {
	return &FullBorder{
		SideBorder: &SideBorder{Border: &Border{display: display}},
	}
}
