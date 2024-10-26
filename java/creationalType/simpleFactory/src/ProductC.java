public class ProductC implements Product {
    private  String name;

    ProductC(){
        this.name = "product C";
    }

    @Override
    public void Say() {
        System.out.println(this.name);
    }
}

