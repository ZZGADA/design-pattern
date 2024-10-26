public class Main {
    public static void main(String[] args) {

        // 实例化工厂
        ProductFactory productFactoryA = new ProductFactoryA();
        ProductFactory productFactoryB = new ProductFactoryB();
        ProductFactory productFactoryC = new ProductFactoryC();

        // 工厂实例化对象
        // product 实例对象 延迟执行
        Product productA = productFactoryA.createProduct();
        Product productB = productFactoryB.createProduct();
        Product productC = productFactoryC.createProduct();

        // 实例的对象执行方法
        productA.say();
        productB.say();
        productC.say();
    }
}