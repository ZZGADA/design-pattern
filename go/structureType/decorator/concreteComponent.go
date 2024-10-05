package decorator

type CoffeeConcreteComponent struct {
}

func (*CoffeeConcreteComponent) GetDescription() string {
	return "这是一杯很普通的咖啡"
}

func (*CoffeeConcreteComponent) GetPrice() float32 {
	return 32.0
}
