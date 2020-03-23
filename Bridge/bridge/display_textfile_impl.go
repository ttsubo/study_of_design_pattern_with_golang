package bridge

import (
	"fmt"
	"os"
)

// DisplayTextfileImpl is struct
type DisplayTextfileImpl struct {
	fileName string
	f        *os.File
}

// NewDisplayTextfileImpl func for initializing DisplayTextfileImpl
func NewDisplayTextfileImpl(str string) *DisplayTextfileImpl {
	return &DisplayTextfileImpl{
		fileName: str,
	}
}

func (d *DisplayTextfileImpl) rawOpen() {
	fp, err := os.Open(d.fileName)
	if err != nil {
		panic(err)
	} else {
		d.f = fp
	}
}

func (d *DisplayTextfileImpl) rawPrint() {
	buf := make([]byte, 64)
	_, err := d.f.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", buf)
}

func (d *DisplayTextfileImpl) rawClose() {
	d.f.Close()
}
