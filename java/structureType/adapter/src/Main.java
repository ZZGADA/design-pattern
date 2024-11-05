public class Main {
    public static void main(String[] args) {

        AdapteeImpl adaptee = new AdapteeImpl();
        AdapterImpl adapter = new AdapterImpl(adaptee);
        adapter.fitRequest();
    }
}