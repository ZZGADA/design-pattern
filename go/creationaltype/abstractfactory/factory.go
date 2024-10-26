package abstractfactory

// AbstractFactory 是一个抽象工厂接口，定义了一个 Create 方法
type AbstractFactory interface {
	Create(productType string) interface{}
}

// FurnitureFactory 是一个具体工厂类，实现了 AbstractFactory 接口
type FurnitureFactory struct{}

func (f FurnitureFactory) Create(productType string) interface{} {
	switch productType {
	case "Chair":
		return Chair{}
	case "Table":
		return Table{}
	default:
		return nil
	}
}

////////////////////////////////////////////////////////////////////////////////////////

// ApplianceFactory 是一个具体工厂类，实现了 AbstractFactory 接口
type ApplianceFactory struct{}

func (a ApplianceFactory) Create(productType string) interface{} {
	switch productType {
	case "Fridge":
		return Fridge{}
	case "Oven":
		return Oven{}
	default:
		return nil
	}
}

////////////////////////////////////////////////////////////////////////////////////////
// 由于go并没有泛型 所以需要一个FactoryProducer来创建工厂 从而隐式掉具体工厂的创建
// 当然这个可以自行选择 不是很重要

// FactoryProducer 是一个工厂生成器类，用于生成具体的工厂
type FactoryProducer struct{}

func (fp FactoryProducer) GetFactory(factoryType string) AbstractFactory {
	switch factoryType {
	case "Furniture":
		return FurnitureFactory{}
	case "Appliance":
		return ApplianceFactory{}
	default:
		return nil
	}
}
