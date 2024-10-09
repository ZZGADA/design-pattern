public class LightOffCommandConcrete implements LightCommand {
    LightReceiver lightReceiver;

    public LightOffCommandConcrete(LightReceiver lightReceiver) {
        this.lightReceiver = lightReceiver;
    }

    @Override
    public void execute() {
        this.lightReceiver.off();
    }
}
