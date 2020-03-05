package main

import (
	"./facade"
)

func startMain() {
	pagemaker := facade.PageMaker{}
	pagemaker.MakeWelcomePage("hyuki@hyuki.com", "welcome1.html")
	pagemaker.MakeWelcomePage("mamoru@hyuki.com", "welcome2.html")
}

func main() {
	startMain()
}
