public class ConcreteHandler2 extends Handler {
    @Override
    public void handleRequest(HandlerContext context, String msg) {
        if (msg.equals("Request2")) {
            System.out.println("ConcreteHandler2 handled the request: " + msg);
        } else if (context.getNextContext() != null) {
            context.getNextContext().proceed(msg);
        }
    }
}