public abstract class IOAbstraction {
    protected ImplementorCPU cpu;

    IOAbstraction(ImplementorCPU cpu){
        this.cpu = cpu;
    }

    abstract void IO(String data);
}
