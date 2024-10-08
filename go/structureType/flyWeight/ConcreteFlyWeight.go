package flyWeight

import "fmt"

type ConcreteFlyWeight interface {
	Say()
}

type ConcreteFlyWeightFirst struct {
}

func (cfw *ConcreteFlyWeightFirst) Say() {
	fmt.Println("cfw first")
}

type ConcreteFlyWeightSecond struct {
}

func (cfw *ConcreteFlyWeightSecond) Say() {
	fmt.Println("cfw second")
}
