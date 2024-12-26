/*
声明一个strategy接口 所有执行策略都要实现这个接口 然后通过策略执行器执行
策略执行器就是一个context上下文struct，详情可见./behaviourType中的strategy部分
strategy：
 1. 单例模式（读写锁🔒）
 2. 延时双删
 3. mq异步更新cache
*/
package main

import (
	strategy "syncache/internel/strategy"
)

var (
	rwLockStrategy            strategy.Strategy
	mqAsyUpdateStrategy       strategy.Strategy
	delayDoubleDeleteStrategy strategy.Strategy
)

func main() {
	// 三种数据库同步redis缓存的形式
	context := strategy.NewContext()
	rwLockStrategy = strategy.NewReadWriteLockStrategy(context)
	delayDoubleDeleteStrategy = strategy.NewDoubleDeleteStrategy(context)

	context.SetStrategy(delayDoubleDeleteStrategy)
	context.Start()
}
