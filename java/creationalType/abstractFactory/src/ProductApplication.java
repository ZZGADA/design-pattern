// 定义一个产品族 家电产品族
public interface ProductApplication {
    void create();
}

// 定义多个产品
class ProductFridge implements ProductApplication {
    @Override
    public void create() {
        System.out.println("Fridge created.");
    }
}

class ProductOven implements ProductApplication {
    @Override
    public void create() {
        System.out.println("Oven created.");
    }
}
