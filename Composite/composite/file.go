package composite

import (
	"fmt"
)

// File is sturct
type File struct {
	name string
	size int
}

// NewFile func for initializing File
func NewFile(name string, size int) *File {
	return &File{
		name: name,
		size: size,
	}
}

func (f *File) getName() string {
	return f.name
}

func (f *File) getSize() int {
	return f.size
}

// Add func for adding file
func (f *File) Add(entry entry) {
	if err := doError(); err != nil {
		fmt.Println(err)
	}
}

// PrintList func for printing directory name
func (f *File) PrintList(prefix string) {
	f.print(prefix)
}

func (f *File) print(prefix string) {
	fmt.Printf("%s/%s (%d)\n", prefix, f.getName(), f.getSize())
}

func doError() error {
	msg := "FileTreatmentException"
	return fmt.Errorf("%s", msg)
}
