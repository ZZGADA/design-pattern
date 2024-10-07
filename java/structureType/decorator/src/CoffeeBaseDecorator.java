/**
 * CoffeeBaseDecorator 抽象的基类装饰器 作为基类 将被具体的扩展子类继承
 */
public abstract class CoffeeBaseDecorator implements CoffeeComponent {
    protected CoffeeComponent coffee;

    public CoffeeBaseDecorator(CoffeeComponent coffee) {
        this.coffee = coffee;
    }

    /**
     * 这里追加了一个字符串 是为了区分让流程变得清楚一点 其实是不需要加这个
     * 因为抽象的基类装饰器 本质是一个封装 方便concrete component 具体装饰器做下一步的增加和扩展
     */
    @Override
    public String getDescription() {
        return coffee.getDescription()+"; 经过抽象装饰器，为coffee加一点牛奶或者巧克力";
    }

    @Override
    public float getCost() {
        return coffee.getCost();
    }
}
