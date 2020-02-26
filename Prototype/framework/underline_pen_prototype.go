package prototype

import (
	"fmt"
	"strings"
)

// UnderlinePen is struct
type UnderlinePen struct {
	Ulchar string
}

// Use func for confirming name
func (u *UnderlinePen) Use(s string) {
	length := len(s)
	line := strings.Repeat(u.Ulchar, (length + 2))
	fmt.Printf("\"%s\"\n", s)
	fmt.Printf("%s\n", line)
	fmt.Println("")
}

func (u *UnderlinePen) createClone() Prototype {
	return &UnderlinePen{u.Ulchar}
}
