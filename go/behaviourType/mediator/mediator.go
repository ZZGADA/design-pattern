package mediator

// ChatMediator 中介者
type ChatMediator interface {
	SendMessage(msg string, colleague *Colleague)
	AddColleague(colleague *Colleague)
}

type ConcreteMediator struct {
	Colleagues []*Colleague
}

func (c *ConcreteMediator) AddColleague(colleague *Colleague) {
	c.Colleagues = append(c.Colleagues, colleague)
}

func (c *ConcreteMediator) SendMessage(msg string, colleague *Colleague) {
	for i := 0; i < len(c.Colleagues); i++ {
		colleagueEach := c.Colleagues[i]
		if colleagueEach != colleague {
			colleagueEach.ReceiveMsg(msg)
		}
	}
}
