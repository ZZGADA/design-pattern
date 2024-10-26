package simplefactory

import "fmt"

type IntityType string

const (
	IntityA IntityType = "IntityA"
	IntityB IntityType = "IntityB"
	IntityC IntityType = "IntityC"
)

type SimpleFactory struct {
}

type BaseStructre struct {
	Name string
}

type BaseInterface interface {
	Say()
}

////////////////////////////////////////////////////////

type StructureA struct {
	BaseStructre
}

func (a *StructureA) Say() {
	fmt.Println(a.Name)
}

////////////////////////////////////////////////////////

type StructureB struct {
	BaseStructre
}

func (b *StructureB) Say() {
	fmt.Println(b.Name)
}

////////////////////////////////////////////////////////

type StructureC struct {
	BaseStructre
}

func (c *StructureC) Say() {
	fmt.Println(c.Name)
}

////////////////////////////////////////////////////////

func (simpleFactory *SimpleFactory) GetIntity(typeIn IntityType) BaseInterface {
	switch typeIn {
	case IntityA:
		return &StructureA{
			BaseStructre: BaseStructre{
				Name: "it is Structre A",
			},
		}
	case IntityB:
		return &StructureB{
			BaseStructre: BaseStructre{
				Name: "it is Structre B",
			},
		}
	case IntityC:
		return &StructureB{
			BaseStructre: BaseStructre{
				Name: "it is Structre C",
			},
		}
	default:
		return nil
	}
}
