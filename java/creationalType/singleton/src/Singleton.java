public class Singleton {
    private static int singleAttribute;


    // 追加类锁
    public static int Get() {
        if (Singleton.singleAttribute == 0) {
            synchronized (Singleton.class) {
                if (Singleton.singleAttribute == 0) {
                    Singleton.singleAttribute = 1;
                }
            }
        }
        return Singleton.singleAttribute;
    }
}
