package bridge

// DisplayFunc is struct
type DisplayFunc struct {
	impl displayImpl
}

// NewDisplayFunc func for initializing DisplayFunc
func NewDisplayFunc(impl displayImpl) *DisplayFunc {
	return &DisplayFunc{
		impl: impl,
	}
}

func (d *DisplayFunc) open() {
	d.impl.rawOpen()
}

func (d *DisplayFunc) printBody() {
	d.impl.rawPrint()
}

func (d *DisplayFunc) close() {
	d.impl.rawClose()
}

//Display  func for displaying string
func (d *DisplayFunc) Display() {
	d.open()
	d.printBody()
	d.close()
}
