public class Director {
    // Director控制builder的执行顺序
    // builder负责具体的执行逻辑

    Builder builder;

    Director(Builder builder) {
        this.builder = builder;
    }

    // 指挥者的指挥构建顺序
    public Product construct() {
        this.builder.init();
        this.builder.buildPart1();
        this.builder.buildPart2();
        return this.builder.getResult();
    }

}
