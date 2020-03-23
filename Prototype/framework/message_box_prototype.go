package prototype

import (
	"fmt"
	"strings"
)

// MessageBox is struct
type MessageBox struct {
	decochar string
}

// NewMessageBox func for initializing MessageBox
func NewMessageBox(decochar string) *MessageBox {
	return &MessageBox{
		decochar: decochar,
	}
}

// Use func for confirming name
func (m *MessageBox) Use(s string) {
	length := len(s)
	line := strings.Repeat(m.decochar, (length + 4))

	fmt.Println(line)
	fmt.Printf("%s %s %s\n", m.decochar, s, m.decochar)
	fmt.Println(line)
	fmt.Println("")
}

func (m *MessageBox) createClone() Prototype {
	return NewMessageBox(m.decochar)
}
