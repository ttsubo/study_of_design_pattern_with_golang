package builder

// Director is struct
type Director struct {
	Builder builderInterface
}

// Construct func for conducting some methods of builder
func (d *Director) Construct() {
	d.Builder.makeTitle("Greeting")
	d.Builder.makeString("From the morning to the afternoon")
	d.Builder.makeItems([]string{"Good morning", "Hello"})
	d.Builder.makeString("In the evening")
	d.Builder.makeItems([]string{"Good evening", "Good night", "Good bye"})
	d.Builder.close()
}
