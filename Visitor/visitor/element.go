package visitor

import "fmt"

// Entry is interface
type Entry interface {
	getName() string
	getSize() int
	Accept(v Visitor)
	toString() string
	getDir() []Entry
}

// File is struct
type File struct {
	Entry
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
func (f *File) Add(entry Entry) {
	if err := doError(); err != nil {
		fmt.Println(err)
	}
}

// Accept func for accepting something
func (f *File) Accept(v Visitor) {
	v.visit(f)
}

func (f *File) toString() string {
	return fmt.Sprintf("%s (%d)", f.getName(), f.getSize())
}

// Directory is sturct
type Directory struct {
	name string
	dir  []Entry
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
	for _, f := range d.dir {
		size += f.getSize()
	}
	return size
}

// Add func for adding directory
func (d *Directory) Add(entry Entry) {
	d.dir = append(d.dir, entry)
}

// Accept func for accepting something
func (d *Directory) Accept(v Visitor) {
	v.visit(d)
}

func (d *Directory) toString() string {
	return fmt.Sprintf("%s (%d)", d.getName(), d.getSize())
}

func (d *Directory) getDir() []Entry {
	return d.dir
}

func doError() error {
	msg := "FileTreatmentException"
	return fmt.Errorf("%s", msg)
}
