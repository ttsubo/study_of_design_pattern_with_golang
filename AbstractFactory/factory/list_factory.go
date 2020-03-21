package factory

import "fmt"

// ListFactory is struct
type ListFactory struct {
}

// CreateLink func for creating Link
func (lf *ListFactory) CreateLink(caption, url string) ItemInterface {
	return newListLink(caption, url)
}

// CreateTray func for creating Tray
func (lf *ListFactory) CreateTray(caption string) TrayInterface {
	return newListTray(caption)
}

// CreatePage func for creating page
func (lf *ListFactory) CreatePage(title, author string) PageInterface {
	return newListPage(title, author)
}

// ListLink is struct
type ListLink struct {
	*Link
}

func newListLink(caption, url string) *ListLink {
	return &ListLink{
		Link: &Link{
			Item: &Item{
				caption: caption,
			},
			url: url,
		},
	}
}

func (ll *ListLink) makeHTML() string {
	return fmt.Sprintf("  <li><a href=\"%s\">%s</a></li>\n", ll.url, ll.caption)
}

// ListTray is struct
type ListTray struct {
	*Tray
}

func newListTray(caption string) *ListTray {
	return &ListTray{
		Tray: &Tray{
			Item: &Item{
				caption: caption,
			},
		},
	}
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
	*Page
}

func newListPage(title, author string) *ListPage {
	return &ListPage{
		Page: &Page{
			title:  title,
			author: author,
		},
	}
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
