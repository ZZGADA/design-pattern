### 观察者模式

### 组成元素

观察者模式：定义对象间的**一种一对**
多依赖关系，使得每当一个对象状态发生改变时，其相关依赖对象皆得到通知并被自动更新。

**Key Point**：观察者模式从发布-订阅的角度来理解比较好，即我们有一个subject主题，当主题的内容更新后向**订阅者**
进行主动推送，并不是一个持续长久监听的模式。
同时注意⚠️：我们一般通过代理的形式向订阅者发送消息，因为订阅者只提供一个消费消息的方法，不需要一个持久的实例。

| 元素               | 名称     | 类型                               |
|------------------|--------|----------------------------------|
| Subject          | 抽象目标接口 | interface 或者 abstract structure  |
| ConcreteSubject  | 具体目标   | structure                        |
| Observer         | 抽象观察者  | interface  或者 abstract structure |
| ConcreteObserver | 具体的观察者 | structure                        |

### 说明

1. 这个模式还是比较简单的，按照我上面的说的`Key Posint`描述来看，我们将观察者模式理解成消息发布者的主动推送，而不要理解成观察者对象的异步监听
2. 那么这个模式的适用场景是什么呢？
    1. 当我们的目标实体对象发生改变的时候，需要向实体的观察者进行消息发送，执行观察者的目标逻辑。例如b站中我们关注的up主发布作品，我们订阅者会收到提醒
        2. 但是在真实场景中我们一般用消息队列。但是设想一想为什么要用**代理对象**。一个up主会有1000万的订阅者，难道我们要为吧这1000万个对象放到内存中吗
   