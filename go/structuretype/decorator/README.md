## 装饰器模式

### 组成元素

装饰器设计模式的组成元素主要有四个

| 元素                 | 名称     | 类型        |
|--------------------|--------|-----------|
| Component          | 抽象组件   | interface |
| concrete component | 具体组件   | class     |
| Decorator          | 抽象装饰器  | interface |
| concrete decorator | 装饰器实体类 | class     |

&nbsp;&nbsp; `其中Component` 是装饰器模式中的顶级接口，`concrete component`、`Decorator`、`concrete decorator` 元素都要实现
`Component` 接口。
⚠️需要注意的是：`concrete decorator`的实现形式是通过继承抽象类`Decorator`的形式实现的。

### 装饰器说明：

装饰器模式是对`concrete component` 一个具体实例的`水平扩展`或者说是`增强`。
我们不改变`concrete component`组件原有的代码逻辑，而是在其原有的逻辑上做一个增加或者扩展。  
熟悉spring的小伙伴就会发现这个概念有点似曾相识。没错，这个概念与spring中的aop概念一致。
在代码中，我们的实例经过一层装饰之后，可以依然通过原有的方法进行调用，但是方法返回的结果却发生了变化。

### go独有的特性

首先我们要知道 go是一个语法简洁的语言，go中没有面向对象语法，但是有内嵌和interface的特性可以帮助我们实现继承和多态。
通过内嵌我们可以调用被内嵌对象的方法和成员变量而不需要再次声明一个成员变量用来显示的绑定，从而我们降低我们代码的耦合程度。
所以我们可以直接通过下民的方式直接调用`concrete component`的方法。  
但是这样写其实也违背了装饰器模式的设计初衷，我们在`concrete decorator`中要屏蔽掉 `concrete component`
的实现细节然后对方法进行`扩展`或者`增强`
所以我们还是建议保留一个`Decorator`用来统一的封装，保证多个`concrete decorator`在调用`concrete component`的一致性

```go
package decorate

type CoffeeComponent interface {
	GetDescription() string
	GetPrice() float32
}

type CoffeeDecoratorAddMilk struct {
	CoffeeComponent
}

func (am *CoffeeDecoratorAddMilk) GetDescription() string {
	return am.CoffeeComponent.GetDescription() + ";加入牛奶"
}

func (am *CoffeeDecoratorAddMilk) GetPrice() float32 {
	return am.CoffeeComponent.GetPrice() + float32(11)
}
```

### 说明

偷懒一下😁～  
关于装饰器和适配器的区别可以移步到[java-decorator-README.md](../../../java/structureType/decorator/README.md)
