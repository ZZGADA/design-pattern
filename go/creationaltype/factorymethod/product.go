package factorymethod

import "fmt"

type Product interface {
	SayName()
}

////////////////////////////////////////////////////////

type ProductA struct {
	name string
}

func (p *ProductA) SayName() {
	fmt.Println(p.name)
}

////////////////////////////////////////////////////////

type ProductB struct {
	name string
}

func (p *ProductB) SayName() {
	fmt.Println(p.name)
}

////////////////////////////////////////////////////////

type ProductC struct {
	name string
}

func (p *ProductC) SayName() {
	fmt.Println(p.name)
}
