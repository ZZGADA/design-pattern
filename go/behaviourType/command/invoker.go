package command

type RemoteInvoker struct {
	lightOnCommand  Command
	lightOffCommand Command

	isOpen bool
}

func NewRemoteInvoker(cOn LightCommand, cOff LightCommand) *RemoteInvoker {
	return &RemoteInvoker{
		lightOnCommand:  cOn,
		lightOffCommand: cOff,
	}
}

func (remoter *RemoteInvoker) PushButton() {
	if remoter.isOpen {
		remoter.lightOnCommand.Execute()
	} else {
		remoter.lightOffCommand.Execute()
	}

	remoter.isOpen = !remoter.isOpen
}
