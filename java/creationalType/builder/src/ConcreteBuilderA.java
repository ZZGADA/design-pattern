public class ConcreteBuilderA extends Builder{
    ProductA product;



    /**
     * 核型的执行逻辑 负责创建不同的product
     */
    @Override
    void init() {
        this.product = new ProductA();
    }

    @Override
    void buildPart1() {
        product.setRadio("radio for A");
    }

    @Override
    void buildPart2() {
        product.setSpeaker("speaker for A");
    }

    @Override
    Product getResult() {
        return this.product;
    }
}
