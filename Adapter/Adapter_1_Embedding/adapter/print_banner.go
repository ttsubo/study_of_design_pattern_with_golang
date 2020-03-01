package adapter

// PrintBanner is struct
type PrintBanner struct {
	*Banner
}

// PrintWeak func for formatting with paren
func (p *PrintBanner) PrintWeak() {
	p.ShowWithParen()
}

// PrintString func for formatting with aster
func (p *PrintBanner) PrintString() {
	p.ShowWithAster()
}

// NewPrintBanner func for creating new PrintBanner
func NewPrintBanner(String string) *PrintBanner {
	return &PrintBanner{
		Banner: &Banner{String: String},
	}
}
