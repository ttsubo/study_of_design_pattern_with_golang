package factoryMethod

// Factory is struct
type Factory struct {
}

// FactoryInterface is interface
type FactoryInterface interface {
	createProduct(owner string) ProductInterface
	registerProduct(ProductInterface)
	Create(factory FactoryInterface, owner string) ProductInterface
}

// Create func for creating product
func (f *Factory) Create(factory FactoryInterface, owner string) ProductInterface {
	p := factory.createProduct(owner)
	factory.registerProduct(p)
	return p
}

// ProductInterface is interface
type ProductInterface interface {
	Use()
	Owner() string
}
