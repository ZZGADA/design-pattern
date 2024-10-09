### 命令模式

### 组成元素

| 元素       | 名称   | 类型        |
|----------|------|-----------|
| Command  | 命令接口 | interface |
| Concrete | 具体命令 | class     |
| Receiver | 接受者  | class     |
| Receiver | 接受者  | class     |

* Command： 顶级的命令接口 所有具体的Command都要实现Command ，Command可能有多个 因为具体的命令执行者，可能可以执行不止一个命令


