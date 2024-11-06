public class Main {
    public static void main(String[] args) {
        ConcreteBuilderA concreteBuilderA = new ConcreteBuilderA();
        ConcreteBuilderB concreteBuilderB = new ConcreteBuilderB();


        Director directorA = new Director(concreteBuilderA);
        Director directorB = new Director(concreteBuilderB);

        Product constructA = directorA.construct();
        Product constructB = directorB.construct();

        constructA.describe();
        constructB.describe();
    }
}