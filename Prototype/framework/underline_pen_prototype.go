package prototype

import (
	"fmt"
	"strings"
)

// UnderlinePen is struct
type UnderlinePen struct {
	ulchar string
}

// NewUnderlinePen func for initializing UnderlinePen
func NewUnderlinePen(ulchar string) *UnderlinePen {
	return &UnderlinePen{
		ulchar: ulchar,
	}
}

// Use func for confirming name
func (u *UnderlinePen) Use(s string) {
	length := len(s)
	line := strings.Repeat(u.ulchar, (length + 2))

	fmt.Printf("\"%s\"\n", s)
	fmt.Println(line)
	fmt.Println("")
}

func (u *UnderlinePen) createClone() Prototype {
	return NewUnderlinePen(u.ulchar)
}
