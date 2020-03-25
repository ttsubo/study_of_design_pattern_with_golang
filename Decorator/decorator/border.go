package decorator

type border struct {
	*display
	neighorDisplay displayInterface
}

// SideBorder is struct
type SideBorder struct {
	*border
	borderChar string
}

// NewSideBorder func for initalizing SideBorder
func NewSideBorder(displayIf displayInterface, borderChar string) *SideBorder {
	sideBorder := &SideBorder{
		border: &border{
			display:        &display{},
			neighorDisplay: displayIf,
		},
		borderChar: borderChar,
	}
	sideBorder.myDisplay = sideBorder
	return sideBorder
}

func (s *SideBorder) getColumns() int {
	return len(s.borderChar)*2 + s.neighorDisplay.getColumns()
}

func (s *SideBorder) getRows() int {
	return s.neighorDisplay.getRows()
}

func (s *SideBorder) getRowText(row int) string {
	return s.borderChar + s.neighorDisplay.getRowText(row) + s.borderChar
}

// FullBorder is struct
type FullBorder struct {
	*SideBorder
}

// NewFullBorder func for initalizing SideBorder
func NewFullBorder(displayIf displayInterface) *FullBorder {
	fullBorder := &FullBorder{
		SideBorder: &SideBorder{
			border: &border{
				display:        &display{},
				neighorDisplay: displayIf}},
	}
	fullBorder.myDisplay = fullBorder
	return fullBorder
}

func (f *FullBorder) getColumns() int {
	return 2 + f.neighorDisplay.getColumns()
}

func (f *FullBorder) getRows() int {
	return 2 + f.neighorDisplay.getRows()
}

func (f *FullBorder) getRowText(row int) string {
	if row == 0 {
		return "+" + f.makeLine("-", f.neighorDisplay.getColumns()) + "+"
	} else if row == f.neighorDisplay.getRows()+1 {
		return "+" + f.makeLine("-", f.neighorDisplay.getColumns()) + "+"
	} else {
		return "|" + f.neighorDisplay.getRowText(row-1) + "|"
	}
}

func (f *FullBorder) makeLine(char string, count int) string {
	buf := ""
	for i := 0; i < count; i++ {
		buf += char
	}
	return buf
}
