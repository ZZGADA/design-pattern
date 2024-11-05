package adapter

import "fmt"

// Adapter 适配器 为了适应 Adaptee
type Adapter interface {
	FitRequest()
}

// AdapterImpl 本地适配器适配Adaptee 目的是让adapter.SpecificRequest()的方法可以适配我们自己的功能
type AdapterImpl struct {
	Adaptee
}

// NewAdapter 内嵌的形式注入 Adaptee对象
func NewAdapter(adaptee Adaptee) Adapter {
	return &AdapterImpl{
		Adaptee: adaptee,
	}
}

// fiRequest 适配操作
func (adapter *AdapterImpl) FitRequest() {
	// 调用 第三方包提供的接口 然后对其功能进行增强和适配
	fmt.Printf("the pointer of adaptee is %p\n", adapter)
	adapter.SpecificRequest()
	fmt.Println("ok, now is adapter is using it")
	fmt.Println("ths function of adaptee has been intensified")
}
