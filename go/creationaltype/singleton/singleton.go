package singleton

import (
	"log"
	"sync"
)

type Singleton struct {
	instance int
	sync.Mutex
}

/*
- Get
单例模式 的获取方法 使用双重检验
1. 优先判断instance这个成员变量是否存在，如果不存在就尝试加锁，否则直接返回
2. 加锁成功后再次判断成员变量是否存在 ，因为存在多协程同时竞争锁资源 然后先后获得锁 所以需要再次判断
*/
func (s *Singleton) Get() int {
	if s.instance == 0 {
		log.Println("没有实例，准备获锁")
		s.Lock()
		if s.instance == 0 {
			log.Println("+++++++++ 成功获锁，成功实例成员变量 +++++++++")
			s.instance = 1
		} else {
			log.Println("成功获锁，但是发现成员变量已经实例了")
		}
		s.Unlock()
	}
	return s.instance
}
