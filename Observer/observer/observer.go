package observer

import (
	"fmt"
	"time"
)

// Observer is interface
type Observer interface {
	update(generator *NumberGenerator)
}

// DigitObserver is struct
type DigitObserver struct {
}

func (d *DigitObserver) update(generator *NumberGenerator) {
	fmt.Printf("DigitObservser: %d\n", generator.getNumber())
	time.Sleep(time.Millisecond * 100)
}

// GraphObserver is struct
type GraphObserver struct {
}

func (g *GraphObserver) update(generator *NumberGenerator) {
	fmt.Printf("GraphicObserver:")
	count := generator.getNumber()
	for i := 0; i < count; i++ {
		fmt.Printf("*")
	}
	fmt.Println("")
	time.Sleep(time.Millisecond * 100)
}
