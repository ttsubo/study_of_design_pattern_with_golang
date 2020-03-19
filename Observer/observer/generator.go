package observer

import (
	"math/rand"
	"time"
)

// NumberGenerator is struct
type NumberGenerator struct {
	number    int
	observers []Observer
}

// AddObserver func for adding Observer
func (n *NumberGenerator) AddObserver(observer Observer) {
	n.observers = append(n.observers, observer)
}

func (n *NumberGenerator) notifyObserver() {
	for _, o := range n.observers {
		o.update(n)
	}
}

func (n *NumberGenerator) getNumber() int {
	return n.number
}

// RandomNumberGenerator is struct
type RandomNumberGenerator struct {
	*NumberGenerator
}

// NewRandomNumberGenerator func for initializing RandomNumberGenerator
func NewRandomNumberGenerator() *RandomNumberGenerator {
	return &RandomNumberGenerator{
		NumberGenerator: &NumberGenerator{number: 0},
	}
}

// Execute func for executing something
func (r *RandomNumberGenerator) Execute() {
	for i := 0; i < 20; i++ {
		rand.Seed(time.Now().UnixNano())
		r.number = rand.Intn(49)
		r.notifyObserver()
	}
}
