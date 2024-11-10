import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class Main {
    public static void main(String[] args) throws InterruptedException {
        // 创建处理器
        Handler handler1 = new ConcreteHandler1();
        Handler handler2 = new ConcreteHandler2();
        Handler handler3 = new ConcreteHandler3();

        // 创建 ExecutorService
        ExecutorService executor = Executors.newFixedThreadPool(10);


        // 发送请求
        for (int i = 0; i < 10; i++) {
            final int requestId = i;
            executor.submit(() -> {
                // 设置上下文
                // 创建 HandlerContext 并设置处理链
                HandlerContext context3 = new HandlerContext(executor, handler3, null);
                HandlerContext context2 = new HandlerContext(executor, handler2, context3);
                HandlerContext context1 = new HandlerContext(executor, handler1, context2);
                HandlerContextHolder.setContext(context1);

                String request = "Request" + (requestId % 4); // 模拟不同的请求
                context1.proceed(request);

                HandlerContextHolder.clearContext();
            });
        }

        Thread.sleep(3*1000);
        // 关闭 ExecutorService
        executor.shutdown();
    }
}