package builder

import "fmt"

// TextBuilder is struct
type TextBuilder struct {
	buffer string
}

// NewTextBuilder func for initializing TextBuilder
func NewTextBuilder() *TextBuilder {
	return &TextBuilder{}
}

func (t *TextBuilder) makeTitle(title string) {
	t.buffer += "======================\n"
	t.buffer += fmt.Sprintf("# %s\n", title)
	t.buffer += "\n"
}

func (t *TextBuilder) makeString(str string) {
	t.buffer += fmt.Sprintf("*** %s ***\n", str)
}

func (t *TextBuilder) makeItems(items []string) {
	for _, i := range items {
		t.buffer += fmt.Sprintf("- %s\n", i)
	}
}

func (t *TextBuilder) close() {
	t.buffer += "======================\n"
}

// GetResult func for fetching all from buffer
func (t *TextBuilder) GetResult() string {
	return t.buffer
}
