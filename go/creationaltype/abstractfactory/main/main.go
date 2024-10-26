package main

import "abstractfactory"

func main() {

	factoryProducer := abstractfactory.FactoryProducer{}

	// 获取家具工厂
	furnitureFactory := factoryProducer.GetFactory("Furniture")
	if furnitureFactory != nil {
		// 创建家具产品 需要断言
		chair := furnitureFactory.Create("Chair").(abstractfactory.Furniture)
		chair.Create()
		table := furnitureFactory.Create("Table").(abstractfactory.Furniture)
		table.Create()
	}

	// 获取家电工厂
	applianceFactory := factoryProducer.GetFactory("Appliance")
	if applianceFactory != nil {
		// 创建家电产品
		fridge := applianceFactory.Create("Fridge").(abstractfactory.Appliance)
		fridge.Create()
		oven := applianceFactory.Create("Oven").(abstractfactory.Appliance)
		oven.Create()
	}

}
