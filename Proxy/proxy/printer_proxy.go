package proxy

import (
	"fmt"
	"time"
)

// Printable in interface
type Printable interface {
	SetPrinterName(name string)
	GetPrinterName() string
	MyPrint(str string)
}

type printer struct {
	name string
}

var instance *printer

func getPrinter(name string) *printer {
	if instance == nil {
		instance = newPrinter(name)
	} else {
		instance.name = name
	}
	return instance
}

func newPrinter(name string) *printer {
	prt := &printer{name: name}
	prt.heavyJob(fmt.Sprintf("Printerのインスタンス(%s)を作成中", name))
	return prt
}

func (p *printer) SetPrinterName(name string) {
	p.name = name
}

func (p *printer) GetPrinterName() string {
	return p.name
}

func (p *printer) MyPrint(str string) {
	fmt.Printf("=== Printer使用者(%s) ===\n", p.name)
	fmt.Println(str)
	fmt.Println("")
}

func (p *printer) heavyJob(msg string) {
	fmt.Printf(msg)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Printf(".")
	}
	fmt.Println("完了")
}

// PrinterProxy is struct
type PrinterProxy struct {
	Name, real string
}

// SetPrinterName func for setting name in printer
func (p *PrinterProxy) SetPrinterName(name string) {
	p.Name = name
}

// GetPrinterName func for fetching name
func (p *PrinterProxy) GetPrinterName() string {
	return p.Name
}

// MyPrint func for printing something
func (p *PrinterProxy) MyPrint(str string) {
	real := getPrinter(p.Name)
	real.MyPrint(str)
}
