package bridge

import (
	"fmt"
)

// RandomIORefined 外层执行器
// IOAbstractionBase 继承抽象父类
type RandomIORefined struct {
	IOAbstractionBase
}

func NewRandomIORefined(cpu ImplementorCPU) IOAbstraction {
	return &RandomIORefined{
		IOAbstractionBase{
			cpu: cpu,
		},
	}
}

// IO 子类实例抽象接口的IO方法 并调用抽象父类的cpu成员变量的Flush方法
func (io *RandomIORefined) IO(data string) {
	fmt.Println("随机IO")
	// 调用父类的抽象方法
	io.cpu.Flush(data)
}
