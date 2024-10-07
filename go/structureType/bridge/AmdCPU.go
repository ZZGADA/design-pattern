package bridge

import "fmt"

type AmdCPU struct {
}

func NewAmdCPU() ImplementorCPU {
	return &AmdCPU{}
}

func (amd *AmdCPU) Flush(data string) {
	fmt.Println("Amd 刷入磁盘" + data)
}
