package builder

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// Construct 控制执行流程
func (d *Director) Construct() Product {
	d.builder.initProduct()
	d.builder.builderPart1()
	d.builder.builderPart2()
	return d.builder.getResult()
}
