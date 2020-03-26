package main

import (
	"./facade"
)

func startMain() {
	facade.Pagemaker.MakeWelcomePage("hyuki@hyuki.com", "welcome1.html")
	facade.Pagemaker.MakeWelcomePage("mamoru@hyuki.com", "welcome2.html")
}

func main() {
	startMain()
}
