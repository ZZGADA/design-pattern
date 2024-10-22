### 单例模式

### 组成元素

单例模式：单例模式确保某一个类只有一个实例，而且自行实例化并向整个系统提供这个实例，这个类称为单例类，它提供全局访问的方法。

| 元素               | 名称    | 类型    |
|------------------|-------|-------|
| Singleton        | 实体类   | class |


### 说明

1. 如果是学过java的小伙伴对这个就更加的熟悉了。这个是一个静态类，其提供一个静态的方法，允许所有线程访问并获得其中的唯一实例
2. 多线程的访问是通过 **“双重校验”**的方法实现的，很简单，很易懂。我们直接上代码
3. PS：这个也是AOP的核型设计模式哦

```java
public class Singleton {
    private static int instance;


    // 追加类锁
    public static int Get() {
        if (Singleton.instance == 0) {
            System.out.println("实例不存在～～");
            synchronized (Singleton.class) {
                if (Singleton.instance == 0) {
                    System.out.println("+++++++++ 获取锁成功 实例化成功 +++++++++");
                    Singleton.instance = 1;
                }else{
                    System.out.println("获取锁成功 但是实例已经存在");
                }
            }
        }
        return Singleton.instance;
    }
}


```