package observer

import "fmt"

type Observer interface {
	Update(data string)
}

type ConcreteObserver struct {
	id string
}

func (c *ConcreteObserver) Update(data string) {
	fmt.Printf("Observer %s: Received data: %s\n", c.id, data)
}
