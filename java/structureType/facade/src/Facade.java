public class Facade {
    ModuleA moduleA;
    ModuleB moduleB;

    Facade(){
        moduleA = new ModuleA();
        moduleB = new ModuleB();
    }

    public boolean useApi() {
        try{
            this.moduleA.test();
            this.moduleB.test();
        }catch (Exception e){
            System.out.println(e.getMessage());
            return false;
        }
        return true;
    }
}
