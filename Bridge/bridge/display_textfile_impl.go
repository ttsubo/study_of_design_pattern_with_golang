package bridge

import (
	"fmt"
	"os"
)

// DisplayTextfileImpl is struct
type DisplayTextfileImpl struct {
	FileName string
	fp       *os.File
}

func (d *DisplayTextfileImpl) rawOpen() {
	f, err := os.Open(d.FileName)
	d.fp = f
	if err != nil {
		panic(err)
	}
}

func (d *DisplayTextfileImpl) rawPrint() {
	buf := make([]byte, 64)
	_, err := d.fp.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", buf)
}

func (d *DisplayTextfileImpl) rawClose() {
	d.fp.Close()
}
