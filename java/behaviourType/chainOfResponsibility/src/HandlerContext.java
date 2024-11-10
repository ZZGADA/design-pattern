import java.util.concurrent.ExecutorService;

public class HandlerContext {
    private final ExecutorService executor;
    private final Handler currentHandler;
    private final HandlerContext nextContext;

    public HandlerContext(ExecutorService executor, Handler currentHandler, HandlerContext nextContext) {
        this.executor = executor;
        this.currentHandler = currentHandler;
        this.nextContext = nextContext;
    }

    public ExecutorService getExecutor() {
        return executor;
    }

    public Handler getCurrentHandler() {
        return currentHandler;
    }

    public HandlerContext getNextContext() {
        return nextContext;
    }

    // process 只处理本handler的逻辑
    public void proceed(String msg) {
        if (currentHandler != null) {
            currentHandler.handleRequest(this, msg);
        }
    }
}