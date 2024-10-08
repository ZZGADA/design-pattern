import java.util.HashMap;
import java.util.Map;

public class FlyWeightFactory {
    private static FlyWeightFactory instance;
    private Map<String, FlyWeight> flyWeight;

    public FlyWeightFactory(HashMap<String, FlyWeight> map) {
        this.flyWeight = map;
    }

    public static FlyWeightFactory GetFlyWeightFactory() {
        if (FlyWeightFactory.instance == null) {
            // 加锁校验
            synchronized (FlyWeightFactory.class) {
                // 双重教研
                if (FlyWeightFactory.instance == null) {
                    FlyWeightFactory.instance = new FlyWeightFactory(new HashMap<>());
                }
            }
        }
        return FlyWeightFactory.instance;
    }

    public FlyWeight get(String key){
        if (!instance.flyWeight.containsKey(key)) {
            synchronized (instance.flyWeight) {
                // 双重验证
                // 保证多线程访问情况下 访问资源时候 只有一个线程可以实例化
                if (!instance.flyWeight.containsKey(key)) {
                    ConcreteFlyWeightSimpleFactory concreteFlyWeightSimpleFactory = ConcreteFlyWeightSimpleFactory.GetFactory();
                    instance.flyWeight.put(key,concreteFlyWeightSimpleFactory.GetConcreteFlyWeight(key));
                }
            }
        }
        return instance.flyWeight.get(key);
    }

}
