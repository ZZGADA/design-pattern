public class ConcreteFlyWeightSecond implements FlyWeight {
    // 每一个flyweight的具体的执行逻辑和操作可能不一样
    // 所以不需要写一个接口方法做为实现
    // 但是按照实际的业务场景 可以改变
    public void say(){
        System.out.println("it is ConcreteFlyWeightSecond function");
    }
}
