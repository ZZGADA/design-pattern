public class ProductA implements Product {
    private  String name;

    ProductA(){
        this.name = "product A";
    }

    @Override
    public void Say() {
        System.out.println(this.name);
    }
}
