package facade

import (
	"errors"
	"fmt"
)

type ModuleA struct{}

func (moduleA ModuleA) test() error {
	fmt.Println("it is module a ")
	return nil
}

type ModuleB struct{}

func (moduleB ModuleB) test() error {
	fmt.Println("it is module b ")
	return nil
}

// Facade 外观模式
type Facade struct {
	moduleA ModuleA
	moduleB ModuleB
}

func NewFacade() *Facade {
	return &Facade{moduleA: ModuleA{}, moduleB: ModuleB{}}
}

// UseApi 外观模式 对外暴露的接口
func (facade *Facade) UseApi() error {
	// 外观对外暴露一个方法，然后将内部的实现细节进行屏蔽
	// 对于调用者来说，内部的实现逻辑是一个黑盒 不需要关系
	// 而内部执行逻辑涉及多个步骤 虽然只依次执行 但是最终的返回结果只有要么全部成功 要么全部失败

	if err := facade.moduleA.test(); err != nil {
		return errors.New(fmt.Sprintf("module a error: %v", err))
	}
	if err := facade.moduleB.test(); err != nil {
		return errors.New(fmt.Sprintf("module b error: %v", err))
	}

	// 上述全部执行成功返回nil
	return nil
}
