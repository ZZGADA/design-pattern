package main

import (
	"decorator"
	"fmt"
)

/*
  - 让我们来看看go的代码
    这里我写了两种写法 原因是利用了go的内嵌特性
    原因让我们看看README
*/
func main() {
	var coffeeComponent decorator.CoffeeComponent
	coffeeComponent = &decorator.CoffeeConcreteComponent{}
	fmt.Println(coffeeComponent.GetDescription())
	fmt.Println(coffeeComponent.GetPrice())

	fmt.Println("------------------------------")

	coffeeComponent = &decorator.CoffeeDecoratorAddMilk{CoffeeComponent: coffeeComponent}
	fmt.Println(coffeeComponent.GetDescription())
	fmt.Println(coffeeComponent.GetPrice())

	fmt.Println("------------------------------")

	chocolate := decorator.CoffeeDecoratorChocolate{CoffeeDecorator: &decorator.CoffeeDecorator{CoffeeComponent: coffeeComponent}}
	fmt.Println(chocolate.GetDescription())
	fmt.Println(chocolate.GetPrice())
}
