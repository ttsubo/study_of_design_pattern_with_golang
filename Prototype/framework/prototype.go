package prototype

// Prototype is interface
type Prototype interface {
	Use(s string)
	createClone() Prototype
}
