package bridge

import (
	"math/rand"
	"time"
)

// DisplayRandomFunc is struct
type DisplayRandomFunc struct {
	*DisplayFunc
}

// NewDisplayRandomFunc func for initializing DisplayRandomFunc
func NewDisplayRandomFunc(impl displayImpl) *DisplayRandomFunc {
	return &DisplayRandomFunc{
		DisplayFunc: &DisplayFunc{
			impl: impl,
		},
	}
}

//RandomDisplay  func for displaying string
func (d *DisplayRandomFunc) RandomDisplay(times int) {
	d.open()
	rand.Seed(time.Now().UnixNano())
	randomTimes := rand.Intn(times)
	for i := 0; i < randomTimes; i++ {
		d.printBody()
	}
	d.close()
}
