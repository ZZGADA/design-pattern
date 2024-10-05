package org.example.sample;

/**
 * CoffeeAddMilkDecorator 具体的装饰类 加入牛奶的装饰类
 */
public class CoffeeAddMilkDecorator extends CoffeeBaseDecorator{
    CoffeeAddMilkDecorator(CoffeeComponent coffee){
        super(coffee);
    }

    /**
     * concrete component 的工作是做扩展和增加 所以使用super的形式 屏蔽掉Decorator 抽象基类的封装过程
     */
    @Override
    public String getDescription() {
        return super.getDescription()+"; 现在加上牛奶";
    }

    @Override
    public float getCost() {
        return super.getCost()+10;
    }
}
