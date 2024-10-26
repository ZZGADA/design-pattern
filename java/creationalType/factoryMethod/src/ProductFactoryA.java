public class ProductFactoryA implements ProductFactory {
    @Override
    public Product createProduct() {
        return new ProductA();
    }
}
