package bridge

// DisplayCountFunc is struct
type DisplayCountFunc struct {
	Impl displayImplInterface
}

func (d *DisplayCountFunc) open() {
	d.Impl.rawOpen()
}

func (d *DisplayCountFunc) printBody() {
	d.Impl.rawPrint()
}

func (d *DisplayCountFunc) close() {
	d.Impl.rawClose()
}

//Display  func for displaying string
func (d *DisplayCountFunc) Display() {
	d.open()
	d.printBody()
	d.close()
}

//MultiDisplay  func for displaying string
func (d *DisplayCountFunc) MultiDisplay(times int) {
	d.open()
	for i := 0; i < times; i++ {
		d.printBody()
	}
	d.close()
}
