package adapter

import "fmt"

type banner interface {
	showWithParen()
	showWithAster()
}

func (p *PrintBanner) showWithParen() {
	fmt.Printf("(%s)\n", p.str)
}

func (p *PrintBanner) showWithAster() {
	fmt.Printf("*%s*\n", p.str)
}
