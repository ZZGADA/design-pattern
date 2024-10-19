package main

import "observer"

func main() {
	subject := observer.NewConcreteSubject()

	subject.Register(observer.NewConcreteObserver("1"))
	subject.Register(observer.NewConcreteObserver("2"))
	subject.Register(observer.NewConcreteObserver("3"))

	subject.Notify("Hello, Observers!")

}
