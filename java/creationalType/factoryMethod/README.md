### 工厂方法模式

### 组成元素

Factory：抽象工厂

* 提供创建公共抽象产品的一个接口

ConcreteFactory: 具体工厂
* 实现抽象工厂的抽象接口，用于创建为一个产品对象
* 对于3种产品就有三个具体工厂用于实现


Product：抽象产品角色

* 抽象产品角色是所创建的所有对象的父类，负责描述所有实例所共有的公共接口

Concrete Product：具体产品角色

* 具体产品角色是创建目标，所有创建的对象都充当这个角色的某个具体类的实例。



| 元素               | 名称      | 类型        |
|------------------|---------|-----------|
| Factory          | 抽象工厂    | interface |
| ConcreteFactory  | 具体工厂    | class     |
| Product          | 抽象产品    | interface |
| Concrete Product | 具体的产品角色 | class     |

### 说明

1. 这个时候可能就会问了，我们只是为了一个实例，有必要先实例这么多工厂吗？  --- 诶别说，还真有点点必要。😁因为工厂的存在帮助我们推延了实例的实例化过程，同时不需要enum或者常量来标识我们多个具体的product，不同的工厂就对应了我们不用的product
* 对于product的实例，正常情况，我们在需要使用其中的方法的时候才会去实例product对吧。现在设想一下，我们不在一个具体的方法实例这个对象，而是在成员变量实例这个product呢
   ```java
   class Test{
      Product product;
      ProductFactory factory;
     
      Test(){
         this.product = new ProductA();
         this.factory = new ProductFactoryA();
      }
   
      public void tellName(){
          Product a = new ProductA();
          a.say();
      }  
     
      public void tellNameUseAttribute(){
          this.product.say();
      }       
   }
  ```
  * 让我们分析一下上面的代码。如果我们是在tellName这个方法里面来实例一个productA，那么当然是合理的。我们都调用了这个方法，那么肯定是执行里面的代码的。
   这样我们在代码中直接实例这个对象就OK了，当然没问题，也不需要什么工厂。 
  * 但是如果我们用的是Test的成员变量 product，那么我们在实例Test的时候这个product对象就存在了，可是这个product对象我们从来没使用过，只有tellNameUseAttribute方法被调用了
   我们才会使用product这个实例，那么我们的内存不就白白浪费了。所以我们可以通过工厂来推延我们product的具体实例化过程，同时我们又绑定了Test 和 ProductA的依赖关系，我们直接用FactoryA的create()方法就可以实例一个productA的对象了
2. 上面这个你看完了，你又会想🤔，那我工厂对象不是还是在Test初始化的时候就加载了，这个内存不是还是浪费了。
  * 这个你说的对，对于这个问题我们没有办法解决，但是我们还需要关注一个东西，那就是实例的大小。我们的ProductFactoryA并没有成员变量，只有一个方法，那么其在编译的时候就会加入到元空间中，不会占用我们宝贵的堆空间，其他什么静态的成员变量都没有
（当然你可以自己自定义的时候添加，但他们都在元空间中）。但是一个ProductA这个实例可能就有很多成员变量了，比如一个数据库对象（BO/DO），这么大一个实例要是在内存中迟迟不释放（无法gc），那么对内存的压力有多大。gc回收不了，分代的话会逐渐进进入老年代。
  ```java
public class ProductFactoryA implements ProductFactory {
    @Override
    public Product createProduct() {
        return new ProductA();
    }
}

```

3. 我们再在上面的基础上做一个拓展，spring默认是单例的，如果不使用工厂方法这种模式，同一个productA实例可能将同时存在多个单例的bean中，然后无法gc，最终变为老年代的。虽然不会造成内存的泄漏但是也白白占用了我们宝贵的内存资源啊。
    如果这个productA的成员变量是一些基础类型倒还行，但是如果维护的是一个TCP连接呢？是一个数据库连接呢？是多个嵌套的其他类呢？
   * 内存在哭泣😭😭😭😭😭
   