package mediator

import "fmt"

type ColleagueMethod interface {
	SendMsg(msg string)
	ReceiveMsg(msg string)
}

type Colleague struct {
	Name     string
	Mediator ChatMediator
}

func NewColleague(name string, mediator ChatMediator) *Colleague {
	return &Colleague{
		Name:     name,
		Mediator: mediator,
	}
}

func (c *Colleague) SendMsg(msg string) {
	fmt.Println(c.Name+"send msg: ", msg)
	c.Mediator.SendMessage(msg, c)
}

func (c *Colleague) ReceiveMsg(msg string) {
	fmt.Println(c.Name+" receive msg: ", msg)
}
