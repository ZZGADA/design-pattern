package observer

// Subject 是被观察者接口
type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	Notify(data string)
}

// ConcreteSubject 是具体的被观察者
type ConcreteSubject struct {
	observers map[Observer]struct{} //	绑定具体的观察者
}

func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{
		observers: make(map[Observer]struct{}),
	}
}

func (cs *ConcreteSubject) Register(observer Observer) {
	/**
	-	在 cs.observers[observer] = struct{}{} 中，
		cs.observers 是一个以 Observer 为键、struct{} 为值的 map。
		使用 struct{} 作为值的类型是因为它不占用额外的内存空间，仅用于表示某个键的存在性。/
	*/
	cs.observers[observer] = struct{}{}
}

func (cs *ConcreteSubject) Deregister(observer Observer) {
	// 去除具体的观察者
	delete(cs.observers, observer)
}

func (cs *ConcreteSubject) Notify(data string) {
	// Notify 通信
	// 遍历所有绑定的观察者 然后对观察者进行通信
	for observer := range cs.observers {
		observer.Update(data)
	}
}
