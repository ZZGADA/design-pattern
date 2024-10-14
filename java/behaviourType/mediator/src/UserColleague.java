public abstract class UserColleague {
    protected ChatMediator mediator;
    protected String name;

    public UserColleague(ChatMediator mediator, String name) {
        this.mediator = mediator;
        this.name = name;
    }

    public abstract void send(String message);
    public abstract void receive(String message);
}
