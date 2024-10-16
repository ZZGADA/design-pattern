public class RealSubject implements Subject {
    @Override
    public void say() {
        System.out.println("it is say function");
    }

    @Override
    public void run() {
        System.out.println("it is run function");
    }
}
