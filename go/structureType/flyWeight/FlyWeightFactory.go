package flyWeight

import "sync"

type FlyWeightFactory struct {
	mapFlyWeight map[string]ConcreteFlyWeight
}

var (
	flyWeightFactoryInstance *FlyWeightFactory
	onceFlyWeightFactory     sync.Once
	lock                     sync.Mutex
)

// GetFlyWeightFactory 获取FlyWeight 工厂
func GetFlyWeightFactory() *FlyWeightFactory {
	onceFlyWeightFactory.Do(func() {
		flyWeightFactoryInstance = &FlyWeightFactory{}
	})
	return flyWeightFactoryInstance
}

func (flyWeightFactory *FlyWeightFactory) Get(key string) ConcreteFlyWeight {
	concreteFwFactory := GetConcreteFlyWeightSimpleFactory()
	_, exists := flyWeightFactory.mapFlyWeight[key]
	if !exists {
		lock.Lock()
		if _, existsDouble := flyWeightFactory.mapFlyWeight[key]; !existsDouble {
			flyWeightFactory.mapFlyWeight[key] = concreteFwFactory.Get(key)
		}
		lock.Unlock()
	}
	return flyWeightFactory.mapFlyWeight[key]
}
