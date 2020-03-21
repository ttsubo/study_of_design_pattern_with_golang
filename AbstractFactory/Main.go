package main

import (
	"flag"

	"./factory"
)

func startMain(factoryObject factory.Factory) {
	asahi := factoryObject.CreateLink("Asahi", "http://www.asahi.com")
	yomiuri := factoryObject.CreateLink("Yomiuri", "http://www.yomiuri.co.jp")
	usYahoo := factoryObject.CreateLink("Yahoo", "http://www.yahoo.com")
	jaYahoo := factoryObject.CreateLink("Yahoo!Japan", "http://www.yahoo.co.jp")
	google := factoryObject.CreateLink("Google", "http://www.google.com")
	excite := factoryObject.CreateLink("Excite", "http://www.excite.co.jp")

	traynews := factoryObject.CreateTray("Newspaper")
	traynews.Add(asahi)
	traynews.Add(yomiuri)

	trayyahoo := factoryObject.CreateTray("Yahoo!")
	trayyahoo.Add(usYahoo)
	trayyahoo.Add(jaYahoo)

	traysearch := factoryObject.CreateTray("Search Engine")
	traysearch.Add(trayyahoo)
	traysearch.Add(excite)
	traysearch.Add(google)

	page := factoryObject.CreatePage("LinkPage", "Hiroshi Yuki")
	page.Add(traynews)
	page.Add(traysearch)
	page.Output(page)
}

func main() {
	flag.Parse()
	if flag.Arg(0) == "ListFactory" {
		startMain(&factory.ListFactory{})
	} else if flag.Arg(0) == "TableFactory" {
		startMain(&factory.TableFactory{})
	}
}
