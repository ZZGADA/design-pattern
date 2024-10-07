public class BaseCoffee implements CoffeeComponent{
    @Override
    public String getDescription() {
        return "这是基础的一杯Coffee";
    }

    @Override
    public float getCost() {
        return 19f;
    }
}
