public class Facade {
    private ModuleA moduleA;
    private ModuleB moduleB;

    Facade() {
        this.moduleA = new ModuleA();
        this.moduleB = new ModuleB();
    }

    public boolean useApi() {
        try {
            this.moduleA.test();
            this.moduleB.test();
        } catch (Exception e) {
            System.out.println(e.getMessage());
            return false;
        }
        return true;
    }
}
