package composite

import "fmt"

// Directory is sturct
type Directory struct {
	Name, path string
	directory  []Entry
}

func (d *Directory) getName() string {
	return d.Name
}

// GetSize func for getting size
func (d *Directory) GetSize() int {
	size := 0
	for _, entry := range d.directory {
		size += entry.GetSize()
	}
	return size
}

// Add func for adding directory
func (d *Directory) Add(entry Entry) {
	d.directory = append(d.directory, entry)
}

// PrintList func for printing directory name
func (d *Directory) PrintList(prefix string) {
	d.print(prefix)
	for _, e := range d.directory {
		e.PrintList(prefix + "/" + d.Name)
	}
}

func (d *Directory) print(prefix string) {
	fmt.Printf("%s/%s (%d)\n", prefix, d.getName(), d.GetSize())
}
