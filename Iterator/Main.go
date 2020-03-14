package main

import (
	"fmt"

	"./iterator"
)

func startMain() {
	bookShelf := iterator.NewBookShelf()
	bookShelf.Append(&iterator.Book{Name: "Aroun d the World in 80 days"})
	bookShelf.Append(&iterator.Book{Name: "Bible"})
	bookShelf.Append(&iterator.Book{Name: "Cinderella"})
	bookShelf.Append(&iterator.Book{Name: "Daddy-Long-Legs"})
	it := bookShelf.Iterator()
	for it.HasNext() {
		book := it.Next()
		fmt.Println(book.GetName())
	}
}

func main() {
	startMain()
}
