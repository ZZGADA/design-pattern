### 外观模式

### 组成元素

外观模式：：外部与一个子系统的通信必须通过一个统一的外观对象进行，为子系统中的一组接口提供一个一致的界面，外观模式定义了一个高层接口，这个接口使得这一子系统更加容易使用。
外观模式又称为门面模式，它是一种对象结构型模式。

**Key Point**：外观模式就是一个大黑盒，我们将我们内部的执行逻辑进行统一的封装。对于外部调用者来说，要们全部成功，要么全部失败。

| 元素        | 名称    | 类型     |
|-----------|-------|--------|
| Facade    | 外观角色  | struct |
| SubSystem | 子系统角色 | struct |

### 说明

1. 外观模式就是做了一层封装📦。如果你是java出身的，那么你一定对封装这个概念非常的熟悉。通过合理的封装，我们屏蔽内部的逻辑细节，外部调用者只管使用就好了。