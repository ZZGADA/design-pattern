package main

import "command"

func main() {
	// 初始化接收者
	receiver := &command.LightReceiver{}

	// 初始化具体命令 然后将命令与具体的接收者绑定
	concreteLightOn := &command.LightOnCommandConcrete{Light: receiver}
	concreteLightOff := &command.LightOffCommandConcrete{Light: receiver}

	// 初始化一个 调用者 用来操作命令
	remoter := command.NewRemoteInvoker(concreteLightOn, concreteLightOff)

	remoter.PushButton() // LightOff
	remoter.PushButton() // LightOn
	remoter.PushButton() // LightOff
	remoter.PushButton() // LightOn

}
