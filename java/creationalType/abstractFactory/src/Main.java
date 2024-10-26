public class Main {
    public static void main(String[] args) {

        // 声明两个抽象工厂
        AbstractFactory<ProductApplication> applicationFactory = new ApplicationFactory();
        AbstractFactory<ProductFurniture> furnitureFactory = new FurnitureFactory();

        // 通过工厂实例产品
        // 创建application 产品族
        ProductApplication fridge = applicationFactory.create(ProductTypeApplication.Fridge.name);
        ProductApplication oven = applicationFactory.create(ProductTypeApplication.Oven.name);


        // 创建furniture 产品族
        ProductFurniture chair = furnitureFactory.create(ProductTypeFurniture.Char.name);
        ProductFurniture table = furnitureFactory.create(ProductTypeFurniture.Table.name);


        // 产品自身的方法
        fridge.create();
        oven.create();

        chair.create();
        table.create();

    }
}