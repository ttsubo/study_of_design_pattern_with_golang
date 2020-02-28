package adapter

// PrintBanner is struct
type PrintBanner struct {
	Banner
	String string
}

// PrintWeak func for formatting with paren
func (p *PrintBanner) PrintWeak() {
	p.Banner.String = p.String
	p.ShowWithParen()
}

// PrintString func for formatting with aster
func (p *PrintBanner) PrintString() {
	p.Banner.String = p.String
	p.ShowWithAster()
}
