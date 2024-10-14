public class PlayingState extends State{
    PlayingState(MediaPlayerContext context){
        super(context);
    }


    public void play() {
        // 当前的状态 如果是playing 就持续播放了 不用管了
        System.out.println("视频已经在播放中了");
    }


    public void pause() {
        // 状态需要切换为停止状态
        System.out.println("暂停播放");
        super.context.switchState(new PauseState(super.context));
    }


    public void stop() {
        // 从播放状态切换到停止状态
        // 需要执行操作 和 更改上下文
        System.out.println("停止播放");
        super.context.switchState(new StopState(super.context));
    }
}
