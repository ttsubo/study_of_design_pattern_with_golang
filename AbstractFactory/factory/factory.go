package factory

import (
	"fmt"
	"os"
)

// Factory is struct
type Factory interface {
	CreateLink(caption, url string) ItemInterface
	CreateTray(caption string) TrayInterface
	CreatePage(title, author string) PageInterface
}

// ItemInterface is interface
type ItemInterface interface {
	makeHTML() string
}

// Item is struct
type Item struct {
	caption string
}

// Link is struct
type Link struct {
	*Item
	url string
}

// TrayInterface is interface
type TrayInterface interface {
	ItemInterface
	Add(item ItemInterface)
}

// Tray is struct
type Tray struct {
	*Item
	tray []ItemInterface
}

// Add func for adding item into Tray
func (t *Tray) Add(item ItemInterface) {
	t.tray = append(t.tray, item)
}

// PageInterface is interface
type PageInterface interface {
	ItemInterface
	Add(item ItemInterface)
	Output(o ItemInterface)
}

// Page is struct
type Page struct {
	title, author string
	content       []ItemInterface
}

// Add func for adding item into Page
func (p *Page) Add(item ItemInterface) {
	p.content = append(p.content, item)
}

// Output func for outputing content
func (p *Page) Output(o ItemInterface) {
	filename := fmt.Sprintf("%s.html", p.title)
	file, _ := os.Create(filename)
	defer file.Close()
	b := []byte(o.makeHTML())
	file.Write(b)
	fmt.Printf("[%s] was created.\n", filename)
}
