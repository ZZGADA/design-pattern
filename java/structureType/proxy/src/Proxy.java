public class Proxy {
    private Subject subject;

    Proxy(String subjectName) {
        // 有的人会说将 真实主题角色作为参数传入
        // 但是我觉得一个代理对象的调用对于客户端就应该是无感知的
        // 所以我就将真实主题角色放在了这里

        switch (subjectName) {
            case "real":
                this.subject = new RealSubject();
                break;
            default:
                this.subject = null;
        }
    }

    // 代理增强的方法
    public void say() {
        this.subject.say();
    }

    // 代理增强的方法
    public void run() {
        this.subject.run();
    }
}
