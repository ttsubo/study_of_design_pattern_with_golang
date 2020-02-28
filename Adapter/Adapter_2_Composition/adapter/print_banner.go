package adapter

// PrintBanner is struct
type PrintBanner struct {
	String string
}

// PrintWeak func for formatting with paren
func (p *PrintBanner) PrintWeak(banner Banner) {
	banner.ShowWithParen()
}

// PrintString func for formatting with aster
func (p *PrintBanner) PrintString(banner Banner) {
	banner.ShowWithAster()
}
