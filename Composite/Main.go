package main

import (
	"fmt"

	"./composite"
)

func startMain() {
	fmt.Println("Making root entries...")
	rootdir := composite.NewDirectory("root")
	bindir := composite.NewDirectory("bin")
	tmpdir := composite.NewDirectory("tmp")
	usrdir := composite.NewDirectory("usr")

	rootdir.Add(bindir)
	rootdir.Add(tmpdir)
	rootdir.Add(usrdir)

	bindir.Add(composite.NewFile("vi", 10000))
	bindir.Add(composite.NewFile("latex", 20000))
	rootdir.PrintList("")

	fmt.Println("")
	fmt.Println("Making user entries...")
	yuki := composite.NewDirectory("yuki")
	hanako := composite.NewDirectory("hanako")
	tomura := composite.NewDirectory("tomura")

	usrdir.Add(yuki)
	usrdir.Add(hanako)
	usrdir.Add(tomura)

	yuki.Add(composite.NewFile("diary.html", 100))
	yuki.Add(composite.NewFile("Composite.java", 200))
	hanako.Add(composite.NewFile("memo.tex", 300))
	tomura.Add(composite.NewFile("game.doc", 400))
	tomura.Add(composite.NewFile("junk.mail", 500))
	rootdir.PrintList("")

	fmt.Println("")
	fmt.Println("Occurring Exception...")
	tmpfile := composite.NewFile("junk.mail", 500)
	bindir = composite.NewDirectory("bin")
	tmpfile.Add(bindir)
}

func main() {
	startMain()
}
