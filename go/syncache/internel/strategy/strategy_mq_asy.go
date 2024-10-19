package strategy

import "fmt"

// MqAsyUpdateStrategy 使用mq异步同步mq
type MqAsyUpdateStrategy struct {
	context Context
}

func (msu *MqAsyUpdateStrategy) run() {
	fmt.Println("mq异步更新")
}
