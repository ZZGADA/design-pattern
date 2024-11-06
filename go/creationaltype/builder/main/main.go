package main

import "builder"

func main() {
	// 初始化builder建造者
	a := &builder.ConcreteBuilderA{}
	b := &builder.ConcreteBuilderB{}

	// 初始化指挥者
	directorA := builder.NewDirector(a)
	directorB := builder.NewDirector(b)

	productA := directorA.Construct()
	productB := directorB.Construct()

	productA.Describe()
	productB.Describe()
}
