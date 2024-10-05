package decorator

type CoffeeDecoratorAddMilk struct {
	CoffeeComponent
}

func (am *CoffeeDecoratorAddMilk) GetDescription() string {
	return am.CoffeeComponent.GetDescription() + ";加入牛奶"
}

func (am *CoffeeDecoratorAddMilk) GetPrice() float32 {
	return am.CoffeeComponent.GetPrice() + float32(11)
}

type CoffeeDecoratorChocolate struct {
	*CoffeeDecorator
}

func (am *CoffeeDecoratorChocolate) GetDescription() string {
	return am.CoffeeDecorator.GetDescription() + ";加入巧克力"
}

func (am *CoffeeDecoratorChocolate) GetPrice() float32 {
	return am.CoffeeDecorator.GetPrice() + float32(22)
}
