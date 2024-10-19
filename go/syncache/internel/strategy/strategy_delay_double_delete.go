package strategy

import "fmt"

// DelayDoubleDeleteStrategy 延时双删策略
type DelayDoubleDeleteStrategy struct {
	context Context
}

func (d *DelayDoubleDeleteStrategy) run() {
	fmt.Println("延时双删策略")

}
