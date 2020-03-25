package composite

type entry interface {
	PrintList(prefix string)
	getSize() int
}
