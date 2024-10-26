public class ProductFactoryC implements ProductFactory {
    @Override
    public Product createProduct() {
        return new ProductC();
    }
}
