

// 抽象工厂
public interface AbstractFactory<T> {
    T create(String type);
}

// 一个具体的产品族的一个工厂
class FurnitureFactory implements AbstractFactory<ProductFurniture> {
    @Override
    public ProductFurniture create(String type) {
        if (ProductTypeFurniture.Char.name.equalsIgnoreCase(type)) {
            return new ProductChair();
        } else if (ProductTypeFurniture.Table.name.equalsIgnoreCase(type)) {
            return new ProductTable();
        }
        return null;
    }
}

class ApplicationFactory implements AbstractFactory<ProductApplication> {
    @Override
    public ProductApplication create(String type) {
        if (ProductTypeApplication.Fridge.name.equalsIgnoreCase(type)) {
            return new ProductFridge();
        } else if (ProductTypeApplication.Oven.name.equalsIgnoreCase(type)) {
            return new ProductOven();
        }
        return null;
    }
}
