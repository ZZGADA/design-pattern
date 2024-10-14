public class Main {
    public static void main(String[] args) {
        Subject concreteSubject = new ConcreteSubject();

        concreteSubject.registerObserver(new ConcreteObserver(1));
        concreteSubject.registerObserver(new ConcreteObserver(2));
        concreteSubject.registerObserver(new ConcreteObserver(3));

        concreteSubject.notifyObservers();
    }
}