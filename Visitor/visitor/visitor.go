package visitor

import "fmt"

// Visitor is interface
type Visitor interface {
	visit(directory Entry)
}

// ListVisitor is struct
type ListVisitor struct {
	currentdir string
}

// NewListVistor func for initializing ListVisitor
func NewListVistor() *ListVisitor {
	return &ListVisitor{
		currentdir: "",
	}
}

func (l *ListVisitor) visit(directory Entry) {
	fmt.Printf("%s/%s\n", l.currentdir, directory.toString())
	if _, ok := directory.(*Directory); ok {
		savedir := l.currentdir
		l.currentdir = fmt.Sprintf("%s/%s", l.currentdir, directory.getName())
		for _, f := range directory.getDir() {
			f.Accept(l)
		}
		l.currentdir = savedir
	}
}
