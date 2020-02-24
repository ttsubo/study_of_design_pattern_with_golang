package factoryMethod

import "fmt"

// IDCardFactory is struct
type IDCardFactory struct {
	Factory
	owners []string
}

func (i *IDCardFactory) createProduct(owner string) ProductInterface {
	fmt.Printf("I'll create %s's card\n", owner)
	return &IDCardProduct{owner}
}

func (i *IDCardFactory) registerProduct(product ProductInterface) {
	i.owners = append(i.owners, product.Owner())
}

// IDCardProduct is struct
type IDCardProduct struct {
	owner string
}

// Use func for using card
func (i *IDCardProduct) Use() {
	fmt.Printf("I'll use %s's card\n", i.owner)
}

// Owner func for confirming owner
func (i *IDCardProduct) Owner() string {
	return i.owner
}
