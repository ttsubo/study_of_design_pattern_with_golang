package composite

import "fmt"

// Directory is sturct
type Directory struct {
	name      string
	directory []entry
}

// NewDirectory func for initializing Directory
func NewDirectory(name string) *Directory {
	return &Directory{
		name: name,
	}
}

func (d *Directory) getName() string {
	return d.name
}

func (d *Directory) getSize() int {
	size := 0
	for _, entry := range d.directory {
		size += entry.getSize()
	}
	return size
}

// Add func for adding directory
func (d *Directory) Add(entry entry) {
	d.directory = append(d.directory, entry)
}

// PrintList func for printing directory name
func (d *Directory) PrintList(prefix string) {
	d.print(prefix)
	for _, e := range d.directory {
		e.PrintList(prefix + "/" + d.name)
	}
}

func (d *Directory) print(prefix string) {
	fmt.Printf("%s/%s (%d)\n", prefix, d.getName(), d.getSize())
}
