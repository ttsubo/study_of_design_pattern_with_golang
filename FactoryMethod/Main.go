package main

import (
	factoryMethod "./framework"
)

func startMain(factoryObject factoryMethod.FactoryInterface) {
	card1 := factoryObject.Create(factoryObject, "Hiroshi Yuki")
	card2 := factoryObject.Create(factoryObject, "Tomura")
	card3 := factoryObject.Create(factoryObject, "Hanako Sato")
	card1.Use()
	card2.Use()
	card3.Use()
}

func main() {
	factory := factoryMethod.IDCardFactory{}
	startMain(&factory)
}
