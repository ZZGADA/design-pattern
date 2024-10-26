package main

import (
	"simplefactory"
)

func main() {
	// 初始化简单工厂
	simpleFactory := simplefactory.SimpleFactory{}

	// 向工厂传入实体的标识
	// 标识可以是 自定义类型（java中就是enum） 也可以是一个字符串 、int、bool
	// 标识不管是什么都可以，只要保证标识和实体类型可以一一对应

	InstanceA := simpleFactory.GetInstance(simplefactory.EntityA)
	InstanceB := simpleFactory.GetInstance(simplefactory.EntityB)
	InstanceC := simpleFactory.GetInstance(simplefactory.EntityC)

	InstanceA.Say()
	InstanceB.Say()
	InstanceC.Say()
}
