package iterator

// Aggregate is interface
type Aggregate interface {
	Append(book *Book)
	Iterator() Iterator
}
