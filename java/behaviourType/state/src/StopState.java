public class StopState extends State{
    StopState(MediaPlayerContext context){
        super(context);
    }

    @Override
    public void play() {
        System.out.println("视频播放中");
        super.context.switchState(new PlayingState(super.context));
    }

    @Override
    public void pause() {
        System.out.println("视频暂停");
        super.context.switchState(new PauseState(super.context));
    }

    @Override
    public void stop() {
        System.out.println("视频已经停止了");
    }
}
