package builder

// Director is struct
type Director struct {
	builder builder
}

// NewDirector func for initializing Director
func NewDirector(builder builder) *Director {
	return &Director{
		builder: builder,
	}
}

// Construct func for conducting some methods of builder
func (d *Director) Construct() {
	d.builder.makeTitle("Greeting")
	d.builder.makeString("From the morning to the afternoon")
	d.builder.makeItems([]string{"Good morning", "Hello"})
	d.builder.makeString("In the evening")
	d.builder.makeItems([]string{"Good evening", "Good night", "Good bye"})
	d.builder.close()
}
