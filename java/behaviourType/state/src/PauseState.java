public class PauseState extends State{

    PauseState(MediaPlayerContext context){
        super(context);
    }

    @Override
    public void play() {
        System.out.println("视频播放中");
        super.context.switchState(new PlayingState(super.context));
    }

    @Override
    public void pause() {
        System.out.println("视频已经在暂停了");
    }

    @Override
    public void stop() {
        System.out.println("视频停止");
        super.context.switchState(new StopState(super.context));
    }
}
