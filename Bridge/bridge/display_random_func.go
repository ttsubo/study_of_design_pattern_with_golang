package bridge

import (
	"math/rand"
	"time"
)

// DisplayRandomFunc is struct
type DisplayRandomFunc struct {
	Impl displayImplInterface
}

func (d *DisplayRandomFunc) open() {
	d.Impl.rawOpen()
}

func (d *DisplayRandomFunc) printBody() {
	d.Impl.rawPrint()
}

func (d *DisplayRandomFunc) close() {
	d.Impl.rawClose()
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
