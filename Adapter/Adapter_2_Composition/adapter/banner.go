package adapter

import "fmt"

// Banner is interface
type Banner interface {
	ShowWithParen()
	ShowWithAster()
}

// ShowWithParen func for displaying String with formatting paren
func (p *PrintBanner) ShowWithParen() {
	fmt.Printf("(%s)\n", p.String)
}

// ShowWithAster func for displaying String with formatting aster
func (p *PrintBanner) ShowWithAster() {
	fmt.Printf("*%s*\n", p.String)
}
