package main

import (
	"fmt"

	"./composite"
)

func startMain() {
	fmt.Println("Making root entries...")
	rootdir := &composite.Directory{Name: "root"}
	bindir := &composite.Directory{Name: "bin"}
	tmpdir := &composite.Directory{Name: "tmp"}
	usrdir := &composite.Directory{Name: "usr"}

	rootdir.Add(bindir)
	rootdir.Add(tmpdir)
	rootdir.Add(usrdir)

	bindir.Add(&composite.File{Name: "vi", Size: 10000})
	bindir.Add(&composite.File{Name: "latex", Size: 20000})
	rootdir.PrintList("")

	fmt.Println("")
	fmt.Println("Making user entries...")
	yuki := &composite.Directory{Name: "yuki"}
	hanako := &composite.Directory{Name: "hanako"}
	tomura := &composite.Directory{Name: "tomura"}

	usrdir.Add(yuki)
	usrdir.Add(hanako)
	usrdir.Add(tomura)

	yuki.Add(&composite.File{Name: "diary.html", Size: 100})
	yuki.Add(&composite.File{Name: "Composite.java", Size: 200})
	hanako.Add(&composite.File{Name: "memo.tex", Size: 300})
	tomura.Add(&composite.File{Name: "game.doc", Size: 400})
	tomura.Add(&composite.File{Name: "junk.mail", Size: 500})
	rootdir.PrintList("")

	fmt.Println("")
	fmt.Println("Occurring Exception...")
	tmpfile := &composite.File{Name: "junk.mail", Size: 500}
	bindir = &composite.Directory{Name: "bin"}
	tmpfile.Add(bindir)
}

func main() {
	startMain()
}
