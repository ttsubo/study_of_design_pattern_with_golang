package adapter

// PrintBanner is struct
type PrintBanner struct {
	str    string
	banner banner
}

// NewPrintBanner func for initializing PrintBanner
func NewPrintBanner(str string) *PrintBanner {
	printBanner := &PrintBanner{
		str: str,
	}
	printBanner.banner = printBanner
	return printBanner
}

// PrintWeak func for formatting with paren
func (p *PrintBanner) PrintWeak() {
	p.banner.showWithParen()
}

// PrintString func for formatting with aster
func (p *PrintBanner) PrintString() {
	p.banner.showWithAster()
}
