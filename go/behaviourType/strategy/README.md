### 策略模式

### 组成元素

策略模式：定义一系列算法，将每一个算法封装起来，并让它们可以相互替换。策略模式让算法独立于使用它的客户而变化，也称为政策模式(
Policy)。

| 元素               | 名称    | 类型        |
|------------------|-------|-----------|
| Context          | 实体类   | struct    |
| Strategy         | 抽象策略  | interface |
| ConcreteStrategy | 具体策略类 | struct    |

### 说明

1. 这里向大家解释一下这个Context，在直译为上下文，但是可能对很多同学来说非常的抽象，所以我们换一个说法。  
   Context理解为**策略的执行器**，我们通过更改策略执行器中的策略，在具体执行方法体不变的情况下更换策略并执行
2. 然后我就通过执行器执行我们的策略就好了，怎么样 是不是超级简单.

```go
package strategy

type Strategy interface {
	Pay(amount int)
}

type Context struct {
   strategy Strategy
}

// SetStrategy 更换策略
func (c *Context) SetStrategy(strategy Strategy) {
   c.strategy = strategy
}

// 执行策略
func (c *Context) ExecuteStrategy(amount int) {
   c.strategy.Pay(amount)
}

```