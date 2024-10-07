public class RandomIORefined extends IOAbstraction{
    RandomIORefined(ImplementorCPU cpu) {
        super(cpu);
    }

    @Override
    void IO(String data) {
        System.out.println("随机IO");
        super.cpu.flush(data);
    }
}
