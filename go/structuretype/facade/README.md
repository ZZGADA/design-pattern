### 外观模式

### 组成元素

外观模式：

**Key Point**：外观模式就是一个大黑盒，我们将我们内部的执行逻辑进行统一的封装。对于外部调用者来说，要们全部成功，要么全部失败。

```go

func (flyWeightFactory *FlyWeightFactory) Get(key string) ConcreteFlyWeight {
concreteFwFactory := GetConcreteFlyWeightSimpleFactory()
_, exists := flyWeightFactory.mapFlyWeight[key]
if !exists {
lock.Lock()
// 按照单例的设计模式，单例都通过双重验证的形式实现单例的创建和获取
if _, existsDouble := flyWeightFactory.mapFlyWeight[key]; !existsDouble {
flyWeightFactory.mapFlyWeight[key] = concreteFwFactory.Get(key)
}
lock.Unlock()
}
return flyWeightFactory.mapFlyWeight[key]
}

```

| 元素     | 名称  | 类型     |
|--------|-----|--------|
| Facade | 结构体 | struct |

### 说明

1. 外观模式就是做了一层封装📦。如果你是java出身的，那么你一定对封装这个概念非常的熟悉。通过合理的封装，我们屏蔽内部的逻辑细节，外部调用者只管使用就好了。