

/**
 * CoffeeAddChocolateDecorator 具体的装饰类 加入巧克力的装饰类
 */
public class CoffeeAddChocolateDecorator extends CoffeeBaseDecorator {
    public CoffeeAddChocolateDecorator(CoffeeComponent coffee) {
        super(coffee);
    }

    /**
     * concrete component 的工作是做扩展和增加 所以使用super的形式 屏蔽掉Decorator 抽象基类的封装过程
     */
    @Override
    public String getDescription() {
        return super.getDescription()+"; 现在加上巧克力";
    }

    @Override
    public float getCost() {
        return super.getCost()+20;
    }
}
