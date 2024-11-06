package state

type Context struct {
	State State
}

func (c *Context) Play() {
	c.State.play()
}

func (c *Context) Pause() {
	c.State.pause()
}

func (c *Context) Stop() {
	c.State.stop()
}
