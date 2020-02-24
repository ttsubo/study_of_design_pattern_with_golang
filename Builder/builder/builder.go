package builder

type builderInterface interface {
	makeTitle(title string)
	makeString(str string)
	makeItems(items []string)
	close()
}
