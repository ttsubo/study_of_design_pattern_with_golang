package flyweight

import (
	"fmt"
	"os"
)

type bigCharFactory struct {
	pool map[string]*bigChar
}

var instance *bigCharFactory

func getInstance() *bigCharFactory {
	if instance == nil {
		instance = &bigCharFactory{}
	}
	return instance
}

func (b *bigCharFactory) getBigChar(charname string) *bigChar {
	var bc *bigChar
	if b.pool == nil {
		b.pool = make(map[string]*bigChar)
	}
	if _, ok := b.pool[charname]; !ok {
		bc = newBigChar(charname)
		b.pool[charname] = bc
	} else {
		bc = b.pool[charname]
	}
	return bc
}

// BigString is struct
type BigString struct {
	bigchars []*bigChar
}

// NewBigString func for initalizing BigString
func NewBigString(str string) *BigString {
	bigStr := &BigString{}
	factory := getInstance()
	for _, s := range str {
		bigStr.bigchars = append(bigStr.bigchars, factory.getBigChar(string(s)))
	}
	return bigStr
}

// Print func for print something
func (b *BigString) Print() {
	for _, bc := range b.bigchars {
		bc.print()
	}
}

type bigChar struct {
	charname string
	fontdata string
}

func newBigChar(charname string) *bigChar {
	char := &bigChar{charname: charname}
	data := make([]byte, 256)
	f, err := os.Open(fmt.Sprintf("big%s.txt", charname))
	if err == nil {
		f.Read(data)
		char.fontdata = string(data)
	} else {
		char.fontdata = charname + "?"
	}
	return char
}

func (b *bigChar) print() {
	fmt.Println(b.fontdata)
}
