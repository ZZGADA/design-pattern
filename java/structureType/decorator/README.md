## 装饰器模式

### 组成元素

装饰器设计模式的组成元素主要有四个

| 元素                 | 名称     | 类型             |
|--------------------|--------|----------------|
| Component          | 抽象组件   | interface      |
| concrete component | 具体组件   | class          |
| Decorator          | 抽象装饰器  | abstract class |
| concrete decorator | 装饰器实体类 | class          |

&nbsp;&nbsp; `其中Component` 是装饰器模式中的顶级接口，`concrete component`、`Decorator`、`concrete decorator` 元素都要实现
`Component` 接口。
⚠️需要注意的是：`concrete decorator`的实现形式是通过继承抽象类`Decorator`的形式实现的。

### 装饰器说明：

装饰器模式是对`concrete component` 一个具体实例的`水平扩展`或者说是`增强`。
我们不改变`concrete component`组件原有的代码逻辑，而是在其原有的逻辑上做一个增加或者扩展。  
熟悉spring的小伙伴就会发现这个概念有点似曾相识。没错，这个概念与spring中的aop概念一致。
在代码中，我们的实例经过一层装饰之后，可以依然通过原有的方法进行调用，但是方法返回的结果却发生了变化。
```java
public static void main(String[] args) {
    CoffeeComponent baseCoffee = new BaseCoffee();

    // step1
    System.out.println(baseCoffee.getDescription());
    System.out.printf("这杯coffee的价格是：%f\n", baseCoffee.getCost());
    System.out.println("------------------------------------------");

    // 注意看baseCoffee这个实例 我
    // 们虽然是new 了一个新的对象出来 但是依然使用原有的 baseCoffee.getDescription()方法来执行业务逻辑
    // step2代码执行的流程和顺序与step1一致 但是结果变了 这个就是装饰器的作用
    
    // step2
    baseCoffee = new CoffeeAddMilkDecorator(baseCoffee);
    System.out.println(baseCoffee.getDescription());
    System.out.printf("这杯coffee的价格是：%f\n", baseCoffee.getCost());
    System.out.println("------------------------------------------");
}
```

### 装饰器和适配器的区别是什么？
这个问题问的很好，当我第一眼看到装饰器的逻辑的时候，我就在想🤔这个不就是适配器的plus版吗。
但是转念一想，我就发现其实不然。 直接看代码：

```java
class A {
    public String say() {
        System.out.println("it is A");
        return "it is a";
    }
}

class B {
    public A a;
    public C c;
    public D d;
    // 构造函数就不写了

    public String sayRedesignForB() {
        System.out.println("我需要使用A提供的say方法");
        String aReply = a.say();
        System.out.println("A的say方法用了，我需要继续执行我自己的逻辑了，当然也会用到A的say()方法返回给我的数据");
    }
}

public static void main(String[] args) {
    A a = new A();
    B b = new B(a);
    b.sayRedesignForB();
}
```
&nbsp;&nbsp; 在这个例子中，B需要做的事情是使用A提供的已经封装好的say()方法，然后再继续执行的逻辑。
这样就通过一个适配器使用上了A的方法。虽然看上去好像也是对A提供的方法做了一个增强或这拓展，但是不然。
我们从两个角度来想想🤔。
1. 代码上：  
   在代码层面，我们其实是需要先执行sayRedesignForB()的一部分代码从而实现我们自己设想的逻辑，也就是第一个System.out.println()。然后再去调用A的say()方法，通过A的say()结果，我们可以继续执行我们的代码了。
   这样来看是不是发现了适配器和装饰器的不同。
   1. **适配器模式**的主要目的是将一个类的接口转换成客户希望的另一个接口。适配器模式使得原本由于接口不兼容而不能一起工作的类可以一起工作。
   2. **装饰器模式**的主要目的是在不修改现有类的情况下，动态地为对象添加新的行为或功能，从而可以在不修改原始类的情况下扩展其功能。
2. 业务场景上：  
   我们通过jdbc与数据库进行交互，但是不同数据库之间调用方式所涉及的协议或者数据格式可能不一样，我们总不可能为不同的数据库设计一套不同的交互interface吧。
一定是统一的Connection、Statement、ResultSet，那么这个时候就需要针对不同的数据库做适配，这样就不会影响我写具体java代码的体验。
对于不同的数据库我们都是先建立连接，然后创建statement对象，然后执行sql语句。  
    但是同时都是将sql语句传给一个数据库，我是用线程池的形式还是单连接的形式，那么就可以用装饰器设计模式进行设计了。

怎么样，现在是不是看出差异了。

