package iterator

// Iterator is interface
type Iterator interface {
	HasNext() bool
	Next() *Book
}
