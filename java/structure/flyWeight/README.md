### 享元模式

### 组成元素

享元模式：运用共享技术有效地支持大量细粒度对象的复用。系统只使用少量的对象，而这些对象都很相似，状态变化很小，可以实现对象的*
*多次复用**。
由于享元模式要求能够共享的对象必须是细粒度对象，因此它又称为轻量级模式，它是一种对象结构型模式。

**Key Point**：享元模式主要用途是整合大量单例模式的实例，这些单例都有一些共有的特性，都可以是一种接口的实例或者继承出来的子类。
我们一般通过一个大的map来维护这些单例，通过统一的一个Get方法对外暴露来获取这些单例。

```java

public FlyWeight get(String key){
    if (!instance.flyWeight.containsKey(key)) {
        synchronized (instance.flyWeight) {
            // 双重验证
            // 保证多goroutine情况下 访问资源时候 只有一个go routine可以实例化
            if (!instance.flyWeight.containsKey(key)) {
                ConcreteFlyWeightSimpleFactory concreteFlyWeightSimpleFactory = ConcreteFlyWeightSimpleFactory.GetFactory();
                instance.flyWeight.put(key,concreteFlyWeightSimpleFactory.GetConcreteFlyWeight(key));
            }
        }
    }
    return instance.flyWeight.get(key);
}

```

| 元素                        | 名称       | 类型                                          |
|---------------------------|----------|---------------------------------------------|
| Flyweight                 | 抽象享元类    | interface 或者 abstract class                 |
| ConcreteFlyweight         | 具体享元     | class                                       |
| UnsharedConcreteFlyweight | 非共享具体享元类 | class (这个可以不用，没有必要在一个工厂中做逻辑上的区分，拆成两个工厂实现最好) |
| FlyweightFactory          | 享元工程类    | class                                       |



### 说明
1. 这个模式还是比较简单的，按照我上面的说的`Key Point`描述来看，就是一个大的map获取全局的单例变量
2. 那么这个模式的适用场景是什么呢？
    1. 我们可以将统一的全局变量按照类型划分，使其可以抽象成享元，从而在获取单例的时候我们可以减少包的调用，屏蔽更多的细节，让代码变的更加简洁
    2. 但是要注意的是这些享元需要实现统一的接口或者继承一个基类，所以如何将享元进行抽象就需要仔细思考，不然后期一旦对抽象接口或者抽象类修改需要变动的地方就会非常多
