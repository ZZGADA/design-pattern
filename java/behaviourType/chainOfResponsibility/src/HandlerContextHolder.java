// HandlerContextHolder.java
public class HandlerContextHolder {
    private static final ThreadLocal<HandlerContext> contextHolder = new ThreadLocal<>();

    public static void setContext(HandlerContext context) {
        contextHolder.set(context);
    }

    public static HandlerContext getContext() {
        return contextHolder.get();
    }

    public static void clearContext() {
        contextHolder.remove();
    }
}
