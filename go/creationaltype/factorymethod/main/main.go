package main

import "factorymethod"

func main() {
	factoryA := factorymethod.ProductFactoryA{}
	factoryB := factorymethod.ProductFactoryB{}
	factoryC := factorymethod.ProductFactoryC{}

	productA := factoryA.CreateProduct()
	productB := factoryB.CreateProduct()
	productC := factoryC.CreateProduct()

	productA.SayName()
	productB.SayName()
	productC.SayName()
}
