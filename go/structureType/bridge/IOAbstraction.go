package bridge

// 将java中的抽象类拆分 一个interface+struct
// IOAbstractionBase 用于子类继承父类的成员变量
// IOAbstraction 是抽象方法 需要子类实现

// IOAbstraction Abstract 抽象接口
// 抽象方法
type IOAbstraction interface {
	IO(data string)
}

// IOAbstractionBase Abstract 抽象父类 基类
type IOAbstractionBase struct {
	cpu ImplementorCPU
}
