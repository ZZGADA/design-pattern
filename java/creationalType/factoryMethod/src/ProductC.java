public class ProductC implements Product {
    private  String name;

    ProductC(){
        this.name = "product C";
    }

    @Override
    public void say() {
        System.out.println(this.name);
    }
}

