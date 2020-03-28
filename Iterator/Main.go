package main

import (
	"fmt"

	"./iterator"
)

func startMain(bookShelf iterator.Aggregate) {
	bookShelf.Append(iterator.NewBook("Aroun d the World in 80 days"))
	bookShelf.Append(iterator.NewBook("Bible"))
	bookShelf.Append(iterator.NewBook("Cinderella"))
	bookShelf.Append(iterator.NewBook("Daddy-Long-Legs"))
	it := bookShelf.Iterator()
	for it.HasNext() {
		book := it.Next()
		fmt.Println(book.GetName())
	}
}

func main() {
	startMain(iterator.NewBookShelf())
}
