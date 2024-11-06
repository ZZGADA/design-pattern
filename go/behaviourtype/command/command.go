package command

type Command interface {
	Execute()
}

type LightCommand interface {
	Execute()
}

type LightOnCommandConcrete struct {
	Light *LightReceiver
}

type LightOffCommandConcrete struct {
	Light *LightReceiver
}

func (c *LightOnCommandConcrete) Execute() {
	c.Light.LightOn()
}

func (c *LightOffCommandConcrete) Execute() {
	c.Light.LightOff()
}
