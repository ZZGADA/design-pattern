public class RemoteInvoker {
    Command commandLightOn;
    Command commandLightOff;
    boolean flag;

    public RemoteInvoker(Command commandLightOn, Command commandLightOff) {
        this.commandLightOn = commandLightOn;
        this.commandLightOff = commandLightOff;
        this.flag = false;
    }

    public void pressButton() {
        if (!flag) {
            // 开灯
            commandLightOn.execute();
        } else {
            commandLightOff.execute();
        }

        flag = !flag;
    }
}
