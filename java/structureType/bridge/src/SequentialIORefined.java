public class SequentialIORefined extends IOAbstraction{
    SequentialIORefined(ImplementorCPU cpu) {
        super(cpu);
    }

    @Override
    void IO(String data) {
        System.out.println("顺序IO");
        super.cpu.flush(data);
    }
}
