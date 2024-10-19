package decorator

// CoffeeDecorator  匿名内嵌 不要在写一个成员变量了
type CoffeeDecorator struct {
	CoffeeComponent
}

func (cd *CoffeeDecorator) GetDescription() string {
	return cd.CoffeeComponent.GetDescription() + ";经过装饰器"
}

func (cd *CoffeeDecorator) GetPrice() float32 {
	return cd.CoffeeComponent.GetPrice()
}
