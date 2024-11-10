public class ConcreteHandler1 extends Handler {
    @Override
    public void handleRequest(HandlerContext context, String msg) {
        if (msg.equals("Request1")) {
            System.out.println("ConcreteHandler1 handled the request: " + msg);
        } else if (context.getNextContext() != null) {
            context.getNextContext().proceed(msg);
        }
    }
}
