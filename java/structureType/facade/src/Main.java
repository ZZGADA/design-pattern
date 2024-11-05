public class Main {
    public static void main(String[] args) {
        Facade facade = new Facade();
        if(facade.useApi()){
            System.out.println("success");
        }else{
            System.out.println("failed");
        }
    }
}