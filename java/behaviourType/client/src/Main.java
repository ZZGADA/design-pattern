public class Main {
    public static void main(String[] args) {

        // 初始化命令的接收者
        // 命令的执行实体
        LightReceiver lightReceiver = new LightReceiver();

        // 将命令绑定一个接收者 生成具体命令
        LightCommand lightOnCommandConcrete = new LightOnCommandConcrete(lightReceiver);
        LightCommand lightOffCommandConcrete = new LightOffCommandConcrete(lightReceiver);

        // 初始化 命令的调用者
        // 命令的调用者 调用具体命令
        // 调用者只关心命令 不关心命令
        RemoteInvoker remoteInvoker = new RemoteInvoker(lightOnCommandConcrete, lightOffCommandConcrete);

        remoteInvoker.pressButton();    // 开灯
        remoteInvoker.pressButton();    // 关灯
        remoteInvoker.pressButton();    // 开灯
    }
}