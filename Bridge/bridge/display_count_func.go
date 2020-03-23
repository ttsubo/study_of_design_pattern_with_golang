package bridge

// DisplayCountFunc is struct
type DisplayCountFunc struct {
	*DisplayFunc
}

// NewDisplayCountFunc func for initializing DisplayCountFunc
func NewDisplayCountFunc(impl displayImpl) *DisplayCountFunc {
	return &DisplayCountFunc{
		DisplayFunc: &DisplayFunc{
			impl: impl,
		},
	}
}

//MultiDisplay  func for displaying string
func (d *DisplayCountFunc) MultiDisplay(times int) {
	d.open()
	for i := 0; i < times; i++ {
		d.printBody()
	}
	d.close()
}
