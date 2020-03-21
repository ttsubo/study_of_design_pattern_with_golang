package builder

type builder interface {
	makeTitle(title string)
	makeString(str string)
	makeItems(items []string)
	close()
}
