package proxy

import "fmt"

// Subject  主题的抽象接口
type Subject interface {
	say()
	run()
}

// RealSubject 主题的实际结构体
type RealSubject struct {
}

func (this *RealSubject) say() {
	fmt.Println("it is real subject say function")
}

func (this *RealSubject) run() {
	fmt.Println("it is real subject run function")
}
