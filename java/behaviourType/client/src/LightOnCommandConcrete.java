public class LightOnCommandConcrete implements LightCommand {
    LightReceiver lightReceiver;

    public LightOnCommandConcrete(LightReceiver lightReceiver) {
        this.lightReceiver = lightReceiver;
    }

    @Override
    public void execute() {
        this.lightReceiver.on();
    }
}
