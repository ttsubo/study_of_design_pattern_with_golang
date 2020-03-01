package composite

import (
	"fmt"
)

// File is sturct
type File struct {
	Name string
	Size int
}

func (f *File) getName() string {
	return f.Name
}

// GetSize func for getting size
func (f *File) GetSize() int {
	return f.Size
}

// Add func for adding file
func (f *File) Add(entry Entry) {
	if err := doError(); err != nil {
		fmt.Println(err)
	}
}

// PrintList func for printing directory name
func (f *File) PrintList(prefix string) {
	f.print(prefix)
}

func (f *File) print(prefix string) {
	fmt.Printf("%s/%s (%d)\n", prefix, f.getName(), f.GetSize())
}

func doError() error {
	msg := "FileTreatmentException"
	return fmt.Errorf("%s", msg)
}
