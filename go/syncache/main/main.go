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
	"time"
)

var (
	mqAsyUpdateStrategy       strategy.Strategy
	rwLockStrategy            strategy.Strategy
	delayDoubleDeleteStrategy strategy.Strategy
)

func main() {
	// 三种数据库同步redis缓存的形式
	context := strategy.NewContext()
	mqAsyUpdateStrategy = &strategy.ReadWriteLockStrategy{Context: context}

	context.SetStrategy(mqAsyUpdateStrategy)
	for i := 0; i < 3; i++ {
		context.Start()
	}

	time.Sleep(10 * time.Second)
}
