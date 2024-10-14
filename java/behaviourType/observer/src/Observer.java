public abstract class Observer {
    public Integer id;

    Observer(Integer id) {
        this.id = id;
    }

    public abstract void update(String data);
}
