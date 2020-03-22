package main

import (
	factoryMethod "./framework"
)

func startMain(factoryObject factoryMethod.FactoryInterface) {
	card1 := factoryObject.Create("Hiroshi Yuki")
	card2 := factoryObject.Create("Tomura")
	card3 := factoryObject.Create("Hanako Sato")
	card1.Use()
	card2.Use()
	card3.Use()
}

func main() {
	startMain(factoryMethod.NewIDCardFactory())
}
