### 单例模式

### 组成元素

单例模式：单例模式确保某一个类只有一个实例，而且自行实例化并向整个系统提供这个实例，这个类称为单例类，它提供全局访问的方法。

| 元素               | 名称    | 类型        |
|------------------|-------|-----------|
| Singleton        | 实体类   | struct    |


### 说明

1. 如果是学过java的小伙伴对这个就更加的熟悉了。这个是一个静态类，其提供一个静态的方法，允许所有线程访问并获得其中的唯一实例
2. 多线程的访问是通过 **“双重校验”**的方法实现的，很简单，很易懂。我们直接上代码
3. 如果你想有更深的体会的话，就来看看这里哦 [redis缓存与数据库一致性场景](../../syncache)。我花了很多时间在go的单例模式中，成功实现了与aop类似的依赖注入

```go
package singleton

import (
	"log"
	"sync"
)

type Singleton struct {
	instance int
	sync.Mutex
}

/* 
    Get
    单例模式 的获取方法 使用双重检验
    1. 优先判断instance这个成员变量是否存在，如果不存在就尝试加锁，否则直接返回
    2. 加锁成功后再次判断成员变量是否存在 ，因为存在多协程同时竞争锁资源 然后先后获得锁 所以需要再次判断
*/
func (s *Singleton) Get() int {
	if s.instance == 0 {
		log.Println("没有实例，准备获锁")
		s.Lock()
		if s.instance == 0 {
			log.Println("+++++++++ 成功获锁，成功实例成员变量 +++++++++")
			s.instance = 1
		} else {
			log.Println("成功获锁，但是发现成员变量已经实例了")
		}
		s.Unlock()
	}
	return s.instance
}

```