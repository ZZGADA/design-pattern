public class AmdCPU implements ImplementorCPU{
    @Override
    public void flush(String data) {
        System.out.println("Amd 刷入磁盘"+data);
    }
}
