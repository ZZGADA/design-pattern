package factorymethod

type EntityType string

const (
	EntityA EntityType = "EntityA"
	EntityB EntityType = "EntityB"
	EntityC EntityType = "EntityC"
)

// ProductFactory 一个产品一个实现类
type ProductFactory interface {
	CreateProduct() Product
}

////////////////////////////////////////////////////////

// ProductFactoryA 产品A
type ProductFactoryA struct{}

func (p *ProductFactoryA) CreateProduct() Product {
	return &ProductA{
		name: "it is A",
	}
}

////////////////////////////////////////////////////////

type ProductFactoryB struct{}

func (p *ProductFactoryB) CreateProduct() Product {
	return &ProductB{
		name: "it is B",
	}
}

////////////////////////////////////////////////////////

type ProductFactoryC struct{}

func (p *ProductFactoryC) CreateProduct() Product {
	return &ProductC{
		name: "it is C",
	}
}
