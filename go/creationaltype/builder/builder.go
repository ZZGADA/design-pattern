package builder

// Builder 定义Builder
type Builder interface {
	builderPart1()
	builderPart2()
	initProduct()
	getResult() Product
}

/**
- 注意⚠️：在go设计语言里面，我们要让代码变得更加优雅和简洁，我就不需要像java一样写set方法了
*/

// ConcreteBuilderA 具体构建者
type ConcreteBuilderA struct {
	product ProductA
}

func (c *ConcreteBuilderA) builderPart1() {
	c.product.radio = " radio A"
}
func (c *ConcreteBuilderA) builderPart2() {
	c.product.speaker = " speaker A"
}
func (c *ConcreteBuilderA) initProduct() {
	c.product = ProductA{}
}
func (c *ConcreteBuilderA) getResult() Product {
	return &c.product
}

// ConcreteBuilderB 具体构建者
type ConcreteBuilderB struct {
	product ProductB
}

func (c *ConcreteBuilderB) builderPart1() {
	c.product.radio = " radio B"
}
func (c *ConcreteBuilderB) builderPart2() {
	c.product.speaker = " speaker B"
}
func (c *ConcreteBuilderB) initProduct() {
	c.product = ProductB{}
}
func (c *ConcreteBuilderB) getResult() Product {
	return &c.product
}
