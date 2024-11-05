package main

import (
	"facade"
)

func main() {

	// 对于调用者来说 我们是不知道内部的执行逻辑和执行顺序的
	// 在go的世界里，由于包的存在，我们可以很好的通过外观模式对外提供公共的方法 同时屏蔽内部细节
	facadeInstance := facade.NewFacade()
	if err := facadeInstance.UseApi(); err != nil {
		panic(err)
	}
}
