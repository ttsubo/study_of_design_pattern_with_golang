package facade

import (
	"fmt"
	"os"
)

// PageMaker is struct
type PageMaker struct {
}

// Pagemaker is variable
var Pagemaker = &PageMaker{}

// MakeWelcomePage func for making welcome page
func (p *PageMaker) MakeWelcomePage(mailaddr, filename string) {
	username := ""
	database := &Database{}
	prop := database.getProperties("maildata")
	for _, person := range prop {
		if person.Mail == mailaddr {
			username = person.Name
		}
	}
	file, _ := os.Create(filename)
	writer := &HTMLWriter{file}
	writer.title(fmt.Sprintf("Welcom to %s's page!", username))
	writer.paragraph("We'll wait for your sending")
	writer.mailto(mailaddr, username)
	writer.close()
	fmt.Printf("%s is created for %s (%s)\n", filename, mailaddr, username)
}
