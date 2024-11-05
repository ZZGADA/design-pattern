package adapter

import "fmt"

// Adaptee 第三方插件或者包提供的开发接口
type Adaptee interface {
	SpecificRequest()
}

type AdapteeImpl struct{}

func (adaptee AdapteeImpl) SpecificRequest() {
	fmt.Printf("the pointer of adaptee is %p\n", &adaptee)
	fmt.Println("it is adaptee specific request,you need to adapt this function for implement your logic")
}
