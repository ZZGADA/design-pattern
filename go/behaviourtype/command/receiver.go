package command

import "fmt"

type LightReceiver struct {
}

func (l *LightReceiver) LightOn() {
	fmt.Println("LightOn")
}

func (l *LightReceiver) LightOff() {
	fmt.Println("LightOff")
}
