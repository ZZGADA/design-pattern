package strategy

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"math/rand"
	"sync"
	"syncache/conf"
	"syncache/internel/client"
	"syncache/internel/models"
	"syncache/internel/service"
	"syncache/internel/service/impl"
	"syncache/internel/template"
	"syncache/internel/utils"
	"time"
)

/*
ReadWriteLockStrategy 读写锁策略
- 适用场景：
 1. 单结点数据库
    1.1. 其实多结点也可以使用，但是要求在更新线程对redis缓存做更新，这样就太消耗资源了
 2. 典型读多写少场景
*/
type ReadWriteLockStrategy struct {
	sync.Once
	sync.Mutex
	Context Context
	BaseStrategy
	redisClient *redis.Client
	final       int

	labelTreeDao        *models.LabelTreeMapper
	labelTreeService    service.LabelTreeService
	lockStrategyService service.StrategyService
}

// NewReadWriteLockStrategy 初始化对象
func NewReadWriteLockStrategy(context Context) *ReadWriteLockStrategy {
	strategy := &ReadWriteLockStrategy{Context: context}
	strategy.init()
	return strategy
}

// init 依赖注入初始化 构造器模式
func (rws *ReadWriteLockStrategy) init() {
	rws.Do(func() {
		log.Println("懒加载执行依赖注入～ ，单例加载")
		rws.redisClient = client.RedisInstance.Get(conf.Dft.Get())

		rws.labelTreeService = impl.NewLabelTreeService()
		rws.lockStrategyService = impl.NewStrategyService()
	})
}

// run strategy 的run方法接口
func (rws *ReadWriteLockStrategy) run() {
	log.Println("单例模式（读写锁实现）协程号：", utils.GetGoroutineID())
	nums := [5]int{200, 456, 777, 588, 371}
	count := 0
	var wg sync.WaitGroup

	// 模拟20000个并发 36s内完成
	// 最终只有5个请求访问到了数据库 请求成功的数量是19968 成功率几乎等于100%  证明效果很好
	// QPS= 555
	for i := 0; i < 20000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if res, err := rws.kernelStrategy(nums[rand.Int31n(5)]); err != nil {
				// 对外抛出的报错 均为redis报错
				log.Println(err.Error())
				fmt.Println("业务繁忙中，请稍等一下～")
			} else {
				fmt.Println(res)
				count++
			}
			fmt.Println("--------------------------------------------------------------------")
		}()
		//if i%100 == 0 {
		//	time.Sleep(time.Millisecond * 80)
		//}
	}
	wg.Wait()
	fmt.Println("final+++++++++ =====> ", rws.final, "========> ", count)
}

/*
kernelStrategy 单例模式+分布式锁实现的核型逻辑

	现在我们的场景是单redis实例、单mysql结点！！！
	由于redis是单线程的，那么我们读写就是分离的，我们可以不用使用读锁来保证数据在读的时候有写操作更新数据
	所以我们只需要使用写锁来保证写操作成功 删除缓存的操作
	同时保证多个读操作同时读到空数据后可以按照序列形式的将数据从db中存入redis缓存中

	注意：我们还需要额外在本地使用锁结构 防止多个实例1s中同时有大量协程同时处于获取锁的状态 降低redis的压力，同时优化速度
*/
func (rws *ReadWriteLockStrategy) kernelStrategy(labelTreeId int) (string, error) {
	labelTreeKey := fmt.Sprintf("%s:%d", template.RedisKeyLabelTree, labelTreeId)

	// 1. 优先获取缓存 读操作不加锁
	val, err := rws.redisClient.Get(context.Background(), labelTreeKey).Result()
	if errors.Is(err, redis.Nil) {
		log.Println("缓存不存在 需要更新缓存 优先本地并发校验")
		val, err = rws.localConcurrentJudge(labelTreeKey, labelTreeId)
	}

	// 1.2. 缓存存在就直接返回
	return val, err
}

// localConcurrentJudge 本地并发校验
func (rws *ReadWriteLockStrategy) localConcurrentJudge(labelTreeKey string, labelTreeId int) (string, error) {
	// 2.1 尝试获本地锁
	if success := rws.tryGetLocalLock(); success {
		log.Println("成功获取到本地锁")
		defer rws.Unlock()

		val, err := rws.redisClient.Get(context.Background(), labelTreeKey).Result()
		if errors.Is(err, redis.Nil) {
			log.Println("第二次尝试获取缓存，仍然没有获取到缓存")
			val, err = rws.remoteConcurrentJudge(labelTreeKey, labelTreeId)
		}

		return val, err
	}

	// 2.2 本地获锁失败 直接返回 表示并发量特别大
	// 因为已经阻塞了30s了再尝试获取一次缓存
	// 在初始流量高的时候 这里是容器获取到缓存数据的
	if val, err := rws.redisClient.Get(context.Background(), labelTreeKey).Result(); len(val) > 1 {
		return val, err
	}

	return "", errors.New("本地并发太高了")
}

// remoteConcurrentJudge 多实例并发校验
func (rws *ReadWriteLockStrategy) remoteConcurrentJudge(labelTreeKey string, labelTreeId int) (string, error) {
	// 3.1.  key不存在就是没数据 此时加锁准备获取数据
	// 此时保证了本地协程没有竞争了 但是多实例场景没有保证 追加分布式锁
	// 最后出函数体的时候追加将锁删掉
	labelTreeLockKey := fmt.Sprintf("%s:lock", labelTreeKey)
	defer func() {
		// 防止key删不掉
		for i := 0; i < 5; i++ {
			if _, err := rws.redisClient.Del(context.Background(), labelTreeLockKey).Result(); err != nil {
				log.Printf("删除key失败, key:%s , error:%#v", labelTreeLockKey, err)
			} else {
				break
			}
		}
	}()

	if lockRedisSuccess, _ := rws.tryGetRedisLock(labelTreeLockKey); lockRedisSuccess {
		// 4.1 获得分布式锁成功 第三次判断是否有缓存
		val, err := rws.redisClient.Get(context.Background(), labelTreeKey).Result()
		if errors.Is(err, redis.Nil) {
			log.Println("成功获取获取分布式 准备更新缓存")
			rws.final++
			val, err = rws.lockStrategyService.PushSpecificLabelTreeInfoById(labelTreeId)
		}
		return val, err
	}

	return "", errors.New("多实例竞争锁超时")
}

// tryGetLocalLock 获取本地锁操作
func (rws *ReadWriteLockStrategy) tryGetLocalLock() bool {
	// 追加计时器 如果阻塞时间超过30s那么就直接拒绝❌
	// 如果获得锁成功就返回true  否则继续尝试
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	for {
		if success := rws.TryLock(); success {
			return true
		}

		select {
		case <-ctx.Done():
			return false
		default:
			time.Sleep(time.Millisecond * 100)
			continue
		}
	}
}

// tryGetRedisLock 尝试获取redis的分布式锁
func (rws *ReadWriteLockStrategy) tryGetRedisLock(labelTreeKey string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	for {
		success, err := rws.redisClient.SetNX(ctx, labelTreeKey, 1, time.Minute*2).Result()
		if err != nil {
			log.Printf("redis 报错了: %#v\n", err)
			return false, err
		}

		if success {
			// 获取锁成功
			return true, nil
		}

		select {
		case <-ctx.Done():
			return false, nil
		default:
			continue
		}
	}

}
