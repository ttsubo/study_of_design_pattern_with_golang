package main

import (
	prototype "./framework"
)

func startMain(managerObject prototype.ManagerInterface) {
	upen := prototype.UnderlinePen{Ulchar: "-"}
	mbox := prototype.MessageBox{Decochar: "*"}
	sbox := prototype.MessageBox{Decochar: "/"}
	managerObject.Register("strong message", &upen)
	managerObject.Register("warning box", &mbox)
	managerObject.Register("slash box", &sbox)

	p1 := managerObject.Create("strong message")
	p2 := managerObject.Create("warning box")
	p3 := managerObject.Create("slash box")
	p1.Use("Hello World")
	p2.Use("Hello World")
	p3.Use("Hello World")
}

func main() {
	manager := prototype.Manager{}
	startMain(&manager)
}
