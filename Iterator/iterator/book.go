package iterator

// Book is struct
type Book struct {
	name string
}

// NewBook func for initializing Book
func NewBook(name string) *Book {
	return &Book{
		name: name,
	}
}

// GetName func for getting bookname
func (b *Book) GetName() string {
	return b.name
}

type bookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func newBookShelfIterator(bookShelf *BookShelf) *bookShelfIterator {
	return &bookShelfIterator{
		bookShelf: bookShelf,
		index:     0,
	}
}

// HasNext func for checking condition
func (b *bookShelfIterator) HasNext() bool {
	if b.index < b.bookShelf.getLength() {
		return true
	}
	return false
}

// Next func for fetching book
func (b *bookShelfIterator) Next() *Book {
	book := b.bookShelf.getBookAt(b.index)
	b.index++
	return book
}

// BookShelf is struct
type BookShelf struct {
	last  int
	books []*Book
}

// NewBookShelf func for initializing NewBookShelf
func NewBookShelf() *BookShelf {
	return &BookShelf{
		last: 0,
	}
}

func (b *BookShelf) getBookAt(index int) *Book {
	return b.books[index]
}

// Append func for adding book
func (b *BookShelf) Append(book *Book) {
	b.books = append(b.books, book)
	b.last++
}

func (b *BookShelf) getLength() int {
	return b.last
}

// Iterator func for iterating BookShelfIterator
func (b *BookShelf) Iterator() Iterator {
	return newBookShelfIterator(b)
}
