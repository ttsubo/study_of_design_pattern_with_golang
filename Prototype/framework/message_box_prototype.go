package prototype

import (
	"fmt"
	"strings"
)

// MessageBox is struct
type MessageBox struct {
	Decochar string
}

// Use func for confirming name
func (m *MessageBox) Use(s string) {
	length := len(s)
	line := strings.Repeat(m.Decochar, (length + 4))
	fmt.Printf("%s\n", line)
	fmt.Printf("%s %s %s\n", m.Decochar, s, m.Decochar)
	fmt.Printf("%s\n", line)
	fmt.Println("")
}

func (m *MessageBox) createClone() Prototype {
	return &MessageBox{m.Decochar}
}
