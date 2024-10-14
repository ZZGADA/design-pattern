import java.util.ArrayList;
import java.util.List;

public class ConcreteMediator implements ChatMediator {
    private List<UserColleague> users = new ArrayList<>();

    @Override
    public void sendMessage(String message, UserColleague user) {
        // 向所有绑定Mediator的用户发送消息
        // 这个类似于观察者模式
        for (UserColleague u : users) {
            if (u != user) {
                u.receive(message);
            }
        }
    }

    @Override
    public void addUser(UserColleague user) {
        users.add(user);
    }
}
