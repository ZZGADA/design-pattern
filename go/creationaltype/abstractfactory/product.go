package abstractfactory

import "fmt"

// Furniture 是一个接口，定义了一个 Create 方法
type Furniture interface {
	Create()
}

// Chair 是 Furniture 接口的一个实现
type Chair struct{}

func (c Chair) Create() {
	fmt.Println("Chair created.")
}

// Table 是 Furniture 接口的一个实现
type Table struct{}

func (t Table) Create() {
	fmt.Println("Table created.")
}

////////////////////////////////////////////////////////////////////////////////////////

// Appliance 是一个接口，定义了一个 Create 方法
type Appliance interface {
	Create()
}

// Fridge 是 Appliance 接口的一个实现
type Fridge struct{}

func (f Fridge) Create() {
	fmt.Println("Fridge created.")
}

// Oven 是 Appliance 接口的一个实现
type Oven struct{}

func (o Oven) Create() {
	fmt.Println("Oven created.")
}
