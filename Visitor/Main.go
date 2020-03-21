package main

import (
	"fmt"

	"./visitor"
)

func startMain() {
	fmt.Println("Making root entries")
	rootdir := visitor.NewDirectory("root")
	bindir := visitor.NewDirectory("bin")
	tmpdir := visitor.NewDirectory("tmp")
	usrdir := visitor.NewDirectory("usr")

	rootdir.Add(bindir)
	rootdir.Add(tmpdir)
	rootdir.Add(usrdir)

	bindir.Add(visitor.NewFile("vi", 10000))
	bindir.Add(visitor.NewFile("latex", 20000))
	rootdir.Accept(visitor.NewListVistor())

	fmt.Println("")

	fmt.Println("Making user entries")
	yuki := visitor.NewDirectory("yuki")
	hanako := visitor.NewDirectory("hanako")
	tomura := visitor.NewDirectory("tomura")

	usrdir.Add(yuki)
	usrdir.Add(hanako)
	usrdir.Add(tomura)

	yuki.Add(visitor.NewFile("diary.html", 100))
	yuki.Add(visitor.NewFile("composite.py", 200))
	hanako.Add(visitor.NewFile("memo.tex", 300))
	tomura.Add(visitor.NewFile("game.doc", 400))
	tomura.Add(visitor.NewFile("junk.mail", 500))
	rootdir.Accept(visitor.NewListVistor())

	fmt.Println("")
	fmt.Println("Occurring Exception...")
	tmpfile := visitor.NewFile("tmp.txt", 100)
	bindir = visitor.NewDirectory("bin")
	tmpfile.Add(bindir)
}

func main() {
	startMain()
}
