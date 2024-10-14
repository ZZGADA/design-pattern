public class MediaPlayerContext {
    private State state;

    /**
     * 上下文的切换操作 用于切换上下的状态 不同的状态对应不同的操作
     * @param state
     */
    public synchronized void switchState(State state) {
        this.state = state;
    }

    /**
     * 所有的状态切换都有上下文来进行操作
     * 所有的状态在上下文中都是公开透明的
     * 注意⚠️：
     * 1. 当前上下文所在状态 是State保存的
     * 2. 所有State的实例 理论来说都需要实例从各自的状态切换其他状态的方法
     * 3. 因为当前用户的状态是不确定的 需要从当前状态切换到另外一个状态
     * 4. 当然特定情况 如果状态的切换是单链路的 即一个状态A切换到状态B必须经过状态C那么就需要进行限定操作
     */

    public void setState(State state){
        this.state = state;
    }

    public void pressPlay() {
        state.play();
    }

    public void pressPause() {
        state.pause();
    }

    public void pressStop(){
        state.stop();
    }
}
