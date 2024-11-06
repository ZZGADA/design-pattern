public class ConcreteBuilderB extends Builder {
    ProductB product;

    /**
     * 核型的执行逻辑 负责创建不同的product
     */

    @Override
    void init() {
        this.product = new ProductB();
    }

    @Override
    void buildPart1() {
        product.setRadio("radio for B");
    }

    @Override
    void buildPart2() {
        product.setSpeaker("speaker for B");
    }

    @Override
    Product getResult() {
        return this.product;
    }
}
