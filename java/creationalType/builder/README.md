### 建造者模式

### 组成元素

Builder：建造者模式   
将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示。

建造者模式是一步一步创建一个复杂的对象，它允许用户只通过指定复杂对象的类型和内容就可以构建它们，用户不需要知道内部的具体构建细节。
建造者模式属于对象创建型模式。根据中文翻译的不同，建造者模式又可以称为生成器模式。

| 元素              | 名称    | 类型             |
|-----------------|-------|----------------|
| Builder         | 抽象建造者 | abstract class |
| ConcreteBuilder | 具体建造者 | class          |
| Director        | 指挥者   | class          |
| Product         | 产品角色  | class          |

### 说明

1. 对于构建者来说我们依然要维护创建型的核心要义---> **延迟具体的对象的创建**。所以建造者模式通过Director 中的 construct方法实现延迟构建
2. 除此之外，构建者模式强调屏蔽具体product的构建细节。构建顺序交由Director确定，构建逻辑由ConcreteBuilder确定
3. 对于产品Product来说，只负责定义产品自身的属性和操作方法，负责对物理世界进行一个抽象。