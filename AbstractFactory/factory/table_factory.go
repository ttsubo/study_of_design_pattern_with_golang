package factory

import "fmt"

// TableFactory is struct
type TableFactory struct {
}

// CreateLink func for creating Link
func (tf *TableFactory) CreateLink(caption, url string) ItemInterface {
	return newTableLink(caption, url)
}

// CreateTray func for creating Tray
func (tf *TableFactory) CreateTray(caption string) TrayInterface {
	return newTableTray(caption)
}

// CreatePage func for creating page
func (tf *TableFactory) CreatePage(title, author string) PageInterface {
	return newTablePage(title, author)
}

// TableLink is struct
type TableLink struct {
	*Link
}

func newTableLink(caption, url string) *TableLink {
	return &TableLink{
		Link: &Link{
			Item: &Item{
				caption: caption,
			},
			url: url,
		},
	}
}

func (tl *TableLink) makeHTML() string {
	return fmt.Sprintf("<td><a href=%s>%s</a></td>\n", tl.url, tl.caption)
}

// TableTray is struct
type TableTray struct {
	*Tray
}

func newTableTray(caption string) *TableTray {
	return &TableTray{
		Tray: &Tray{
			Item: &Item{
				caption: caption,
			},
		},
	}
}

func (tt *TableTray) makeHTML() string {
	buf := "<td>\n"
	buf += "<table width=\"100%\" border=\"1\"><tr>\n"
	buf += fmt.Sprintf("<td bgcolor=\"#cccccc\" algin=\"center\" colsapn=\"%d\"><b>%s</b></td>\n", len(tt.tray), tt.caption)
	buf += "</tr>\n"
	buf += "<tr>\n"
	for _, item := range tt.tray {
		buf += item.makeHTML()
	}
	buf += "</tr></table>\n"
	buf += "</td>\n"
	return buf
}

// TablePage is struct
type TablePage struct {
	*Page
}

func newTablePage(title, author string) *TablePage {
	return &TablePage{
		Page: &Page{
			title:  title,
			author: author,
		},
	}
}

func (tp *TablePage) makeHTML() string {
	buf := "<html>\n"
	buf += fmt.Sprintf("  <head><title>%s</title></head>\n", tp.title)
	buf += "<body>\n"
	buf += fmt.Sprintf("<h1>%s</h1>", tp.title)
	buf += "<table width=\"80%\" border=\"3\">\n"
	for _, item := range tp.content {
		buf += fmt.Sprintf("<tr>%s</tr>", item.makeHTML())
	}
	buf += "</table>"
	buf += fmt.Sprintf("<hr><adress>%s</adress>", tp.author)
	buf += "</body>\n</html>\n"
	return buf
}
