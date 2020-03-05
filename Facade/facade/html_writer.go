package facade

import (
	"fmt"
	"os"
)

// HTMLWriter is struct
type HTMLWriter struct {
	writer *os.File
}

func (h *HTMLWriter) title(title string) {
	h.writer.Write([]byte("<html>\n"))
	h.writer.Write([]byte("<head>"))
	h.writer.Write([]byte(fmt.Sprintf("<title>%s</title>", title)))
	h.writer.Write([]byte("</head>\n"))
	h.writer.Write([]byte("<body>\n"))
	h.writer.Write([]byte(fmt.Sprintf("<h1>%s</h1>\n", title)))
}

func (h *HTMLWriter) paragraph(msg string) {
	h.writer.Write([]byte(fmt.Sprintf("<p>%s</p>\n", msg)))
}

func (h *HTMLWriter) link(href, caption string) {
	h.writer.Write([]byte(fmt.Sprintf("<a href=\"%s\">%s</a>", href, caption)))
}

func (h *HTMLWriter) mailto(mailaddr, username string) {
	h.link(fmt.Sprintf("mailto:%s", mailaddr), username)
}

func (h *HTMLWriter) close() {
	h.writer.Write([]byte("</body>\n"))
	h.writer.Write([]byte("</html>\n"))
	h.writer.Close()
}
