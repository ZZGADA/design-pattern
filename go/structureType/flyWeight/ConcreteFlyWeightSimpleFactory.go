package flyWeight

import "sync"

var (
	concreteFlyWeightSimpleFactory *ConcreteFlyWeightSimpleFactory
	onceConcreteFlyWeightFactory   sync.Once
)

type ConcreteFlyWeightSimpleFactory struct{}

// GetConcreteFlyWeightSimpleFactory 单例模式创建唯一实例
func GetConcreteFlyWeightSimpleFactory() *ConcreteFlyWeightSimpleFactory {
	onceConcreteFlyWeightFactory.Do(func() {
		concreteFlyWeightSimpleFactory = &ConcreteFlyWeightSimpleFactory{}
	})
	return concreteFlyWeightSimpleFactory
}

func (concreteFlyWeightSimpleFactory *ConcreteFlyWeightSimpleFactory) Get(key string) ConcreteFlyWeight {
	switch key {
	case "ConcreteFlyWeightFirst":
		return &ConcreteFlyWeightFirst{}
	case "ConcreteFlyWeightSecond":
		return &ConcreteFlyWeightSecond{}
	default:
		return nil
	}
}
