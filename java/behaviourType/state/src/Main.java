public class Main {
    public static void main(String[] args) {

        MediaPlayerContext mediaPlayerContext = new MediaPlayerContext();
        PlayingState playingState = new PlayingState(mediaPlayerContext);

        // 初始化
        mediaPlayerContext.setState(playingState);

        mediaPlayerContext.pressPlay();
        mediaPlayerContext.pressPlay();
        mediaPlayerContext.pressPlay();
        mediaPlayerContext.pressPause();

        System.out.println("--------------------------");
        mediaPlayerContext.pressPause();
        mediaPlayerContext.pressPause();
        mediaPlayerContext.pressStop();

        System.out.println("--------------------------");
        mediaPlayerContext.pressStop();
        mediaPlayerContext.pressStop();
        mediaPlayerContext.pressPlay();

        System.out.println("--------------------------");
        mediaPlayerContext.pressPlay();
        mediaPlayerContext.pressPlay();
        mediaPlayerContext.pressStop();

        System.out.println("--------------------------");
    }
}