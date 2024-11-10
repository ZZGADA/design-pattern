public class ConcreteHandler3 extends Handler {
    @Override
    public void handleRequest(HandlerContext context, String msg) {
        if (msg.equals("Request3")) {
            System.out.println("ConcreteHandler3 handled the request: " + msg);
        } else if (context.getNextContext() != null) {
            context.getNextContext().proceed(msg);
        }else{
            System.out.println("ConcreteHandler3-0 handled the request: " + msg);
        }
    }
}
