package builder

import (
	"fmt"
	"os"
)

// HTMLBuilder is struct
type HTMLBuilder struct {
	buffer, filename string
	f                *os.File
	makeTitleCalled  bool
}

// NewHTMLBuilder func for initializing HTMLBuilder
func NewHTMLBuilder() *HTMLBuilder {
	return &HTMLBuilder{}
}

func (h *HTMLBuilder) makeTitle(title string) {
	h.filename = fmt.Sprintf("%s.html", title)
	h.f, _ = os.Create(h.filename)
	h.f.Write([]byte(fmt.Sprintf("<html><head><title>%s</title></head></html>", title)))
	h.f.Write([]byte(fmt.Sprintf("<h1>%s</h1>", title)))
	h.makeTitleCalled = true
}

func (h *HTMLBuilder) makeString(str string) {
	if !h.makeTitleCalled {
		os.Exit(1)
	}
	h.f.Write([]byte(fmt.Sprintf("<p>%s</p>", str)))
}

func (h *HTMLBuilder) makeItems(items []string) {
	if !h.makeTitleCalled {
		os.Exit(1)
	}
	h.f.Write([]byte("<ul>"))
	for _, i := range items {
		h.f.Write([]byte(fmt.Sprintf("<li>%s</li>", i)))
	}
	h.f.Write([]byte("</ul>"))
}

func (h *HTMLBuilder) close() {
	if !h.makeTitleCalled {
		os.Exit(1)
	}
	h.f.Write([]byte("</body></html>"))
	h.f.Close()
}

// GetResult func for replying filename
func (h *HTMLBuilder) GetResult() string {
	return h.filename
}
