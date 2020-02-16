package factory

import "fmt"

// TableFactory is struct
type TableFactory struct {
}

// CreateLink func for creating Link
func (tf *TableFactory) CreateLink(caption, url string) LinkInterface {
	tablelink := TableLink{}
	tablelink.caption = caption
	tablelink.url = url
	return &tablelink
}

// CreateTray func for creating Tray
func (tf *TableFactory) CreateTray(caption string) TrayInterface {
	tableTray := TableTray{}
	tableTray.caption = caption
	return &tableTray
}

// CreatePage func for creating page
func (tf *TableFactory) CreatePage(title, author string) PageInterface {
	tablePage := TablePage{}
	tablePage.title = title
	tablePage.author = author
	return &tablePage
}

// TableLink is struct
type TableLink struct {
	Link
}

func (tl *TableLink) makeHTML() string {
	return fmt.Sprintf("<td><a href=%s>%s</a></td>\n", tl.url, tl.caption)
}

// TableTray is struct
type TableTray struct {
	Tray
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
	Page
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
