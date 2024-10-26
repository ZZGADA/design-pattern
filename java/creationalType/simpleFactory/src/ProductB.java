public class ProductB implements Product {
    private  String name;

    ProductB(){
        this.name = "product B";
    }

    @Override
    public void Say() {
        System.out.println(this.name);
    }
}
