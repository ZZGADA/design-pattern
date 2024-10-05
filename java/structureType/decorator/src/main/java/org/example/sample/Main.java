package org.example.sample;

public class Main {

    /**
     * 装饰器模式
     * 组成元素：
     * 1、 Component            抽象组件      interface
     * 2、 concrete component   具体组件      class
     * 3、 Decorator            抽象装饰器    abstract class
     * 4、 concrete decorator   装饰器实体类  class
     */
    public static void main(String[] args) {
        CoffeeComponent baseCoffee = new BaseCoffee();

        System.out.println(baseCoffee.getDescription());
        System.out.printf("这杯coffee的价格是：%f\n", baseCoffee.getCost());
        System.out.println("------------------------------------------");

        baseCoffee = new CoffeeAddMilkDecorator(baseCoffee);
        System.out.println(baseCoffee.getDescription());
        System.out.printf("这杯coffee的价格是：%f\n", baseCoffee.getCost());
        System.out.println("------------------------------------------");

//        对于baseCoffee 一次只用一个装饰器 如果连续使用装饰器的话 大家可以看看会发生情况
//        原因：我们使用了super 所以装饰器会依次调用父类的方法
//
//        baseCoffee = new CoffeeAddChocolateDecorator(baseCoffee);
//        System.out.println(baseCoffee.getDescription());
//        System.out.printf("这杯coffee的价格是：%f\n", baseCoffee.getCost());
//        System.out.println("------------------------------------------");

    }
}
