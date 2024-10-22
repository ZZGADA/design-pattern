public class Main {
    public static void main(String[] args) throws InterruptedException {
        Thread[] threads = new Thread[100];
        for (int i = 0; i < 100; i++) {
            threads[i] = new Thread() {
                @Override
                public void run() {
                    Singleton.Get();
                }
            };
            threads[i].start();
        }

        for (Thread thread : threads) {
            thread.join();
        }
        System.out.println("end");

    }
}