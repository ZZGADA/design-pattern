## 桥接器模式

### 组成元素

桥接器模式：将抽象部分与它的实现部分进行分离，使它们都可以独立地变化。它是一种对象结构型模式，又称为柄体(Handle and Body)
模式或接口(Interface)模式。

| 元素                  | 名称           | 类型                              |
|---------------------|--------------|---------------------------------|
| Abstraction         | 抽象类          | interface + struct (java中一个抽象类) |
| RefinedAbstraction  | 扩充抽象类（外层执行器） | struct                          |
| ImplementorCPU      | 实现类接口        | interface                       |
| ConcreteImplementor | 具体实现类        | struct                          |

我知道这里你看不懂，但是我做一个具体的解释 你就懂啦。

下文表示的`抽象类`都是指一个`interface`+`struct` struct用于存放成员变量、interface用于实现抽象接口

首先我们更换一个定义，RefinedAbstraction大众解释为扩充抽象类，但是你看的懂吗？
我第一眼是真看不懂，所以我们换一个说法，将RefinedAbstraction称之为一个`外层执行器`。
这个`外层执行器`有多种实例，每个执行器又操作多种“具体实现类”，从而生成一个新的组合`实例`用于执行某一个特定的操作（`实例`
的定义很重要，我最后讲）。

举个例子：我现在有两个cpu 一个是Intel，一个是Amd，其表示的是两个`具体的实现类`。
然后cpu里面将数据从内存写入磁盘的操作，有顺序IO和随机IO两种执行策略，这个称为`执行器`。
这个时候针对一个写操作，Intel 进行顺序IO、Intel进行随机IO、Amd进行顺序IO、Amd进行随机IO，就有四种执行结果。
这个时候我们就会有四种组合出来的`实例`，我们可以对这个四个`实例`有不同的操作。

怎么样这样你是不是就清晰了，但是我猜你还有一个问题，按照这样的设计模式，为什么`Abstraction` 要定义为一个抽象类，而不是一个接口。
原因有如下几点：

1. 如果使用接口的话，每一个接口的实现类也就是`RefinedAbstraction`就必须有自定义一个成员变量用来存放`ConcreteImplementor`
   对象
    1. 如果每一个`RefinedAbstraction`实现类都有一个成员变量为什么不在父类抽象出一个成员变量对吧
2. 这个时候你会想，“那我Abstraction接口里面写的抽象方法的行参传入一个`ImplementorCPU`的实例不就好了”。
    1. 是的这个，这个想法很美好，我们可以少一个成员变量，但是这样的话，我们每一次执行器操作不都要传一个实例，这得多慢啊。
3. 话不多说，我们来看看，一个错误🙅❌的写法。（大家看完赶紧忘掉）

```go
package _brideg

import (
	"fmt"
)

type ImplementorCPU interface {
	Flush(data string)
}

type IntelCPU struct {
}

func (intel *IntelCPU) Flush(data string) {
	fmt.Println("Intel 刷入磁盘" + data)
}

//////////////////////////////////////////////////////////////////////////

type IOAbstraction interface {
	IO(data string)
}

type SequentialIORefined struct {
}

// IO 子类实例抽象接口的IO方法 并调用抽象父类的cpu成员变量的Flush方法
func (seqIo *SequentialIORefined) IO(cpu *IntelCPU, data string) {
	fmt.Println("顺序IO")
	cpu.Flush(data)
}

func main() {
	cpu := &IntelCPU{}
	seqIo := &SequentialIORefined{}
	
	seqIo.IO("执行")
	
	// 这样看是不是就发现每一次执行IO都要传入一个cpu的实例 这样来看是不是就显得很臃肿
	// 而且按照逻辑上来说，我总不可能每执行一次IO就要换一个CPU吧
	//// 所以啊我们就该是在SequentialIORefined的父类里面写一个成员变量

}


```

4. 这个时候你有一个想法，为啥我们不针对CPU也就是`ConcreteImplementor`实现多个接口`Abstraction`,这样也可以实现我们的效果
    1. 最开始我直觉也是这样想的，但是发现不对。这样的话我们每一种IO类型都要有一个接口，10种IO就要有10个`Abstraction`接口。
       这个开销无所谓，因为按照桥接器模式我们也要有10个`RefinedAbstraction`执行器，关键在于`ConcreteImplementor`这个实现类。
       这个实现类将有10个不同方法名用于区分不同的`Abstraction`接口，这样在函数调用上面，我就会有10个if做函数的调用和区分，这显然更加的麻烦。
    2. 所以才会有桥接器设计模式，用于连接**执行器**和**具体实现类**
5. 来看看代码吧，😁嘻嘻





