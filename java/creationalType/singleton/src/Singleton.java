public class Singleton {
    private static int instance;


    // 追加类锁
    public static int Get() {
        if (Singleton.instance == 0) {
            System.out.println("实例不存在～～");
            synchronized (Singleton.class) {
                if (Singleton.instance == 0) {
                    System.out.println("+++++++++ 获取锁成功 实例化成功 +++++++++");
                    Singleton.instance = 1;
                }else{
                    System.out.println("获取锁成功 但是实例已经存在");
                }
            }
        }
        return Singleton.instance;
    }
}
