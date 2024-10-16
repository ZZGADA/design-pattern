public class Main {
    public static void main(String[] args) {
        // 创建代理对象
        Proxy proxy = new Proxy("real");
        proxy.say();
        proxy.run();
    }
}