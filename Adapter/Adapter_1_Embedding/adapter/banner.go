package adapter

import "fmt"

type banner struct {
	str string
}

func (b *banner) showWithParen() {
	fmt.Printf("(%s)\n", b.str)
}

func (b *banner) showWithAster() {
	fmt.Printf("*%s*\n", b.str)
}
