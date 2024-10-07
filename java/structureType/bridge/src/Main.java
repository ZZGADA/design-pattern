public class Main {
    public static void main(String[] args) {

        ImplementorCPU intelCPU = new IntelCPU();
        AmdCPU amdCPU = new AmdCPU();

        SequentialIORefined seqIntel = new SequentialIORefined(intelCPU);
        SequentialIORefined amd = new SequentialIORefined(amdCPU);

        RandomIORefined randIntel = new RandomIORefined(intelCPU);
        RandomIORefined randAmd = new RandomIORefined(amdCPU);


        // 大家重点看 ioAbstractions[i].IO("执行");
        // 同一个方法 调用的结果不一样  （实例不同）
        // 本质上是对IOAbstraction和ImplementorCPU 的一个全排列
        IOAbstraction[] ioAbstractions = new IOAbstraction[]{seqIntel,amd,randIntel,randAmd};
        for (int i = 0; i < ioAbstractions.length; i++) {
            ioAbstractions[i].IO("执行");
        }
    }
}