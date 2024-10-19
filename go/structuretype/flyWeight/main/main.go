package main

import "flyWeight"

func main() {
	factory := flyWeight.GetFlyWeightFactory()

	factory.Get("ConcreteFlyWeightFirst").Say()
	factory.Get("ConcreteFlyWeightSecond").Say()

}
