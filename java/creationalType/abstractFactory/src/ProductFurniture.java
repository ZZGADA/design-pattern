// 定一个产品族  家具产品族
public interface ProductFurniture {
    void create();
}

// 产品族下面有多个具体的产品
class ProductChair implements ProductFurniture {
    @Override
    public void create() {
        System.out.println("Chair created.");
    }
}

class ProductTable implements ProductFurniture {
    @Override
    public void create() {
        System.out.println("Table created.");
    }
}
