package bridge

import "fmt"

type IntelCPU struct {
}

func NewIntelCPU() ImplementorCPU {
	return &IntelCPU{}
}

func (intel *IntelCPU) Flush(data string) {
	fmt.Println("Intel 刷入磁盘" + data)
}
