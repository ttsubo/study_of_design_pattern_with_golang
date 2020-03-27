package main

import (
	"fmt"

	"./proxy"
)

func startMain(p proxy.Printable) {
	fmt.Printf("Printer代理人の名前は現在(%s)です\n", p.GetPrinterName())
	p.MyPrint("Nice to meet you")
	p.SetPrinterName("Bob")
	fmt.Printf("Printer代理人の名前は現在(%s)です\n", p.GetPrinterName())
	p.MyPrint("Hello, World")
}

func main() {
	startMain(proxy.NewPrinterProxy("Alice"))
}
