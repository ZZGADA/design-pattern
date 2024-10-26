public class Main {
    public static void main(String[] args) {

        SimpleFactory simpleFactory = new SimpleFactory();

        Product productA = simpleFactory.getProduct(ProductType.TypeA);
        Product productB = simpleFactory.getProduct(ProductType.TypeB);
        Product productC = simpleFactory.getProduct(ProductType.TypeC);

        productA.Say();
        productB.Say();
        productC.Say();
    }
}