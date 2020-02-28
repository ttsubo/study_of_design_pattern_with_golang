package adapter

import "fmt"

// Banner is struct
type Banner struct {
	String string
}

// ShowWithParen func for displaying String with formatting paren
func (b *Banner) ShowWithParen() {
	fmt.Printf("(%s)\n", b.String)
}

// ShowWithAster func for displaying String with formatting aster
func (b *Banner) ShowWithAster() {
	fmt.Printf("*%s*\n", b.String)
}
