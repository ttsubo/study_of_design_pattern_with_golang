package factory

import "fmt"

// ListFactory is struct
type ListFactory struct {
}

// CreateLink func for creating Link
func (lf *ListFactory) CreateLink(caption, url string) LinkInterface {
	listlink := ListLink{}
	listlink.caption = caption
	listlink.url = url
	return &listlink
}

// CreateTray func for creating Tray
func (lf *ListFactory) CreateTray(caption string) TrayInterface {
	listTray := ListTray{}
	listTray.caption = caption
	return &listTray
}

// CreatePage func for creating page
func (lf *ListFactory) CreatePage(title, author string) PageInterface {
	listPage := ListPage{}
	listPage.title = title
	listPage.author = author
	return &listPage
}

// ListLink is struct
type ListLink struct {
	Link
}

func (ll *ListLink) makeHTML() string {
	return fmt.Sprintf("  <li><a href=\"%s\">%s</a></li>\n", ll.url, ll.caption)
}

// ListTray is struct
type ListTray struct {
	Tray
}

func (lt *ListTray) makeHTML() string {
	buf := "<li>\n"
	buf += fmt.Sprintf("%s\n", lt.caption)
	buf += "<ul>\n"
	for _, item := range lt.tray {
		buf += item.makeHTML()
	}
	buf += "</ul>\n"
	buf += "</li>\n"
	return buf
}

// ListPage is struct
type ListPage struct {
	Page
}

func (lp *ListPage) makeHTML() string {
	buf := "<html>\n"
	buf += fmt.Sprintf("  <head><title>%s</title></head>\n", lp.title)
	buf += "<body>\n"
	buf += fmt.Sprintf("<h1>%s</h1>", lp.title)
	buf += "<ul>"
	for _, item := range lp.content {
		buf += item.makeHTML()
	}
	buf += "</ul>"
	buf += fmt.Sprintf("<hr><adress>%s</adress>", lp.author)
	buf += "</body>\n</html>\n"
	return buf
}
