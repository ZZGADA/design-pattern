public class SimpleFactory {

    public Product getProduct(ProductType type) {
        switch (type) {
            case TypeA: {
                return new ProductA();
            }
            case TypeB: {
                return new ProductB();
            }
            case TypeC: {
                return new ProductC();
            }
            default: {
                return null;
            }
        }
    }
}
