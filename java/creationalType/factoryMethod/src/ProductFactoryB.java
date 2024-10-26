public class ProductFactoryB implements ProductFactory {
    @Override
    public Product createProduct() {
        return new ProductB();
    }
}
