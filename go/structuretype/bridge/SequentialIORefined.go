package bridge

import "fmt"

// SequentialIORefined 外层执行器
// IOAbstractionBase 继承抽象父类
type SequentialIORefined struct {
	IOAbstractionBase
}

func NewSequentialIORefined(cpu ImplementorCPU) IOAbstraction {
	return &SequentialIORefined{IOAbstractionBase{
		cpu: cpu,
	}}
}

// IO 子类实例抽象接口的IO方法 并调用抽象父类的cpu成员变量的Flush方法
func (seqIo *SequentialIORefined) IO(data string) {
	fmt.Println("顺序IO")
	seqIo.cpu.Flush(data)
}
