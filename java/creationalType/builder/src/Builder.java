public abstract class Builder {
    // 抽象父类提供 抽象构造函数
    // 交由具体的concrete builder构建

    // 两个需要构造的模块 构造顺序按照具体的product来定
    abstract void buildPart1();
    abstract void buildPart2();
    abstract Product getResult();
    abstract void init();
}
