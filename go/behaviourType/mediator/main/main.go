package main

import "mediator"

func main() {
	concreteMediator := mediator.ConcreteMediator{Colleagues: make([]*mediator.Colleague, 0)}

	colleague1 := mediator.NewColleague("ZZGEDA", &concreteMediator)
	colleague2 := mediator.NewColleague("ZZGEDA1_1", &concreteMediator)
	colleague3 := mediator.NewColleague("ZZGEDA2_2", &concreteMediator)

	concreteMediator.AddColleague(colleague1)
	concreteMediator.AddColleague(colleague2)
	concreteMediator.AddColleague(colleague3)

	colleague1.SendMsg("Hello guys this is ZZGEDA")
}
