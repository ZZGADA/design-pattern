// 状态接口
public abstract class State {
    // 实现确定上下文的多个状态
    // 然后针对这些状态做上下文的切换

    State(MediaPlayerContext context){
        this.context = context;
    }

    protected MediaPlayerContext context;


    abstract void play();
    abstract void pause();
    abstract void stop();
}
