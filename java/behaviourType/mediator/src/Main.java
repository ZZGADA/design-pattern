public class Main {
    public static void main(String[] args) {
        // 创建多个ConcreteColleague 然后每个ConcreteColleague都绑定ConcreteMediator中介者

        ChatMediator chatRoom = new ConcreteMediator();

        // 传入chatRoom 主要是绑定用户和mediator的关系
        UserColleague user1 = new ConcreteColleague(chatRoom, "Alice");
        UserColleague user2 = new ConcreteColleague(chatRoom, "Bob");
        UserColleague user3 = new ConcreteColleague(chatRoom, "Charlie");

        // 用户注册到mediator
        chatRoom.addUser(user1);
        chatRoom.addUser(user2);
        chatRoom.addUser(user3);

        // 用户发送消息 有mediator这个中间者发送消息
        // 实际是this.mediator.sendMessage(message, this);
        user1.send("Hello, everyone!");
    }
}