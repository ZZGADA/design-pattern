public class AdapterImpl  extends  Target implements Adapter{

    // 调用父类的构造函数
    AdapterImpl(Adaptee adaptee) {
        super(adaptee);
    }

    @Override
    public void fitRequest() {
        // 调用适配者的方法 以适应本地执行逻辑
        this.adaptee.specificRequest();

        System.out.println("ok, now is adapter is using it");
        System.out.println("ths function of adaptee has been intensified");
    }
}
