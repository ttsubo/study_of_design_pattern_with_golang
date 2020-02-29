package bridge

// DisplayFunc is struct
type DisplayFunc struct {
	Impl displayImplInterface
}

func (d DisplayFunc) open() {
	d.Impl.rawOpen()
}

func (d DisplayFunc) printBody() {
	d.Impl.rawPrint()
}

func (d DisplayFunc) close() {
	d.Impl.rawClose()
}

//Display  func for displaying string
func (d *DisplayFunc) Display() {
	d.open()
	d.printBody()
	d.close()
}
