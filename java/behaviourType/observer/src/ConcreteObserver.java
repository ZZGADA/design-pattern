public class ConcreteObserver extends Observer {
    ConcreteObserver(Integer id){
        super(id);
    }

    @Override
    public void update(String data) {
        System.out.printf("it is observe %d , data is %s \n",this.id,data);
    }
}
