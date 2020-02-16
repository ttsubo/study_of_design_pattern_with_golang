package factory

import (
	"fmt"
	"os"
)

// Factory is struct
type Factory interface {
	CreateLink(caption, url string) LinkInterface
	CreateTray(caption string) TrayInterface
	CreatePage(title, author string) PageInterface
}

// Item is struct
type Item struct {
	caption string
}

// ItemInterface is interface
type ItemInterface interface {
	makeHTML() string
}

// Link is struct
type Link struct {
	Item
	url string
}

// LinkInterface is interface
type LinkInterface interface {
	ItemInterface
}

// Tray is struct
type Tray struct {
	Item
	tray []ItemInterface
}

// TrayInterface is interface
type TrayInterface interface {
	ItemInterface
	Add(item ItemInterface)
}

// Add func for adding item into Tray
func (t *Tray) Add(item ItemInterface) {
	t.tray = append(t.tray, item)
}

// Page is struct
type Page struct {
	title, author string
	content       []ItemInterface
}

// PageInterface is interface
type PageInterface interface {
	TrayInterface
	Output(o outputInterface)
}

type outputInterface interface {
	makeHTML() string
}

// Add func for adding item into Page
func (p *Page) Add(item ItemInterface) {
	p.content = append(p.content, item)
}

// Output func for outputing content
func (p *Page) Output(o outputInterface) {
	filename := fmt.Sprintf("%s.html", p.title)
	file, _ := os.Create(filename)
	defer file.Close()
	b := []byte(o.makeHTML())
	file.Write(b)
	fmt.Printf("[%s] was created.\n", filename)
}
