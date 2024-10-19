package strategy

import (
	"syncache/conf"
)

type Context struct {
	config   conf.Config
	strategy Strategy
}

func NewContext() Context {
	return Context{
		config: conf.Dft.Get(),
	}
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) Start() {
	c.strategy.run()
}
