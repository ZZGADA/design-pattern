public class ProductB implements Product {
    private  String name;

    ProductB(){
        this.name = "product B";
    }

    @Override
    public void say() {
        System.out.println(this.name);
    }
}
