package strategy

type Context struct {
	strategy strategy
}

// SetStrategy 更换策略
func (c *Context) SetStrategy(strategy strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(amount int) {
	c.strategy.Pay(amount)
}
