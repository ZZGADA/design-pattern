public class IntelCPU implements ImplementorCPU {
    @Override
    public void flush(String data) {
        System.out.println("Intel 刷入磁盘"+data);
    }
}
