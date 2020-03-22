package factoryMethod

import "fmt"

// IDCardFactory is struct
type IDCardFactory struct {
	*Factory
	owners []string
}

// NewIDCardFactory func for initializing IDCardFactory
func NewIDCardFactory() *IDCardFactory {
	idCardFactory := &IDCardFactory{
		Factory: &Factory{},
	}
	idCardFactory.factory = idCardFactory
	return idCardFactory
}

func (i *IDCardFactory) createProduct(owner string) Product {
	return newIDCardProduct(owner)
}

func (i *IDCardFactory) registerProduct(product Product) {
	i.owners = append(i.owners, product.getOwner())
}

type iDCardProduct struct {
	owner string
}

func newIDCardProduct(owner string) *iDCardProduct {
	fmt.Printf("I'll create %s's card\n", owner)
	return &iDCardProduct{owner}
}

// Use func for using card
func (i *iDCardProduct) Use() {
	fmt.Printf("I'll use %s's card\n", i.owner)
}

func (i *iDCardProduct) getOwner() string {
	return i.owner
}
