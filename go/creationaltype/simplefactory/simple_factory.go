package simplefactory

import "fmt"

type EntityType string

const (
	EntityA EntityType = "EntityA"
	EntityB EntityType = "EntityB"
	EntityC EntityType = "EntityC"
)

type SimpleFactory struct {
}

type BaseStructure struct {
	Name string
}

type BaseInterface interface {
	Say()
}

////////////////////////////////////////////////////////

type StructureA struct {
	BaseStructure
}

func (a *StructureA) Say() {
	fmt.Println(a.Name)
}

////////////////////////////////////////////////////////

type StructureB struct {
	BaseStructure
}

func (b *StructureB) Say() {
	fmt.Println(b.Name)
}

////////////////////////////////////////////////////////

type StructureC struct {
	BaseStructure
}

func (c *StructureC) Say() {
	fmt.Println(c.Name)
}

////////////////////////////////////////////////////////

func (simpleFactory *SimpleFactory) GetInstance(typeIn EntityType) BaseInterface {
	switch typeIn {
	case EntityA:
		return &StructureA{
			BaseStructure: BaseStructure{
				Name: "it is structure A",
			},
		}
	case EntityB:
		return &StructureB{
			BaseStructure: BaseStructure{
				Name: "it is structure B",
			},
		}
	case EntityC:
		return &StructureB{
			BaseStructure: BaseStructure{
				Name: "it is structure C",
			},
		}
	default:
		return nil
	}
}
