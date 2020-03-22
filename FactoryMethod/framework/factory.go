package factoryMethod

// FactoryInterface is interface
type FactoryInterface interface {
	createProduct(owner string) Product
	registerProduct(Product)
	Create(owner string) Product
}

// Factory is struct
type Factory struct {
	factory FactoryInterface
}

// Create func for creating product
func (f *Factory) Create(owner string) Product {
	p := f.factory.createProduct(owner)
	f.factory.registerProduct(p)
	return p
}

// Product is interface
type Product interface {
	Use()
	getOwner() string
}
