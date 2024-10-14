import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public class ConcreteSubject implements Subject {
    // 注意这里⚠️：我们一般通过代理的形式在notifyObservers方法中实例对象然后销毁
    // 我们的订阅者 并不需要一个只持久的对象
    private Map<Integer, Observer> map = new HashMap<>();

    /**
     * 增加synchronized同步锁 保证多线程的、增加Observer是线程安全的
     *
     * @param observer
     */
    @Override
    public synchronized void registerObserver(Observer observer) {
        if(!map.containsKey(observer.id)){
            map.put(observer.id,observer);
        }
    }

    @Override
    public void removeObserver(Observer observer) {

    }

    @Override
    public void notifyObservers() {
        map.forEach((k,v)->{
            v.update("hello observer");
        });
    }
}
