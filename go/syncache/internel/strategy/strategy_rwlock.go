package strategy

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
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
- ReadWriteLockStrategy 读写锁策略
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
	redisClient  *redis.Client
	final        int
	localLockNum int32
	successQuery int32

	labelTreeDao     *models.LabelTreeMapper
	labelTreeService service.LabelTreeService
	strategyService  service.StrategyService
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
		rws.strategyService = impl.NewStrategyService()
	})
}

/*
  - run strategy 的run方法接口
    模拟20000个并发 36s内完成
    最终只有5个请求访问到了数据库 请求成功的数量是19968 成功率几乎等于100%  证明效果很好
    QPS= 555

    随机一个时间执行update 最后查看缓存数据和数据库是否一致
*/
func (rws *ReadWriteLockStrategy) run() {
	log.Println("单例模式（读写锁实现）协程号：", utils.GetGoroutineID())
	maxIteratorNum := 20000
	updateFlag := int(rand.Int31n(int32(maxIteratorNum)))
	nums := [2]int{40, 42}
	newParentLabelTreeId := 34
	var wg sync.WaitGroup

	for i := 0; i < maxIteratorNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := rws.kernelStrategyQuery(nums[rand.Int31n(2)]); err != nil {
				// 对外抛出的报错 均为redis报错
				log.Println(err)
			} else {
				atomic.AddInt32(&rws.successQuery, 1)
			}
		}()
		if i == updateFlag {
			time.Sleep(time.Millisecond * 400)
			if err := rws.kernelStrategyUpdate(40, newParentLabelTreeId); err != nil {
				log.Println(err)
			}
		}
	}
	wg.Wait()
	fmt.Println("query into DB： ", rws.final, ", successful query:  ", rws.successQuery, ", into local lock num: ", rws.localLockNum)
}

/*
  - kernelStrategyUpdate 单例模式+分布式锁 实现的更新逻辑
    场景：读多写少

    思路：
    1. 更新数据的时候 先更新数据库，后删除缓存
    2. 如果是更新体系的话 要删除两个部分的key
    2.1. 体系没有变动之前的 labelTreeId下的子体系  ==> 因为当前node结点改变了 其父级也会改变
    2.2. 体系变动之后 labelTreeId的新子体系 ==> 这些子体系的父级增加了层级 全部都要变
    3. 直接走redis的分布式锁即可

    注意⚠️：
    1. 更新操作本地不用锁，从业务上来说就是并发更新的，而且写操作本身就少，完全不用加本地锁，加锁强行变成有序的更新就让MVCC机制完全失去了作用
    2. 更新操作用的分布式锁和读操作的分布式锁是一个锁，主要作用就是读入操作写缓存和更新操作删缓存有序进行，即单例的更新cache
*/
func (rws *ReadWriteLockStrategy) kernelStrategyUpdate(labelTreeId, newParentLabelTreeId int) error {
	// 查出没有更新前 当前labelTreeId下的子结点
	labelTreeKey := fmt.Sprintf("%s:%d", template.RedisKeyLabelTree, labelTreeId)
	labelTreeLockKey := fmt.Sprintf("%s:lock", labelTreeKey)

	// 获锁成功需要执行两个步骤
	// 1. 更新本地数据库
	// 2. 删除缓存
	if lockRedisSuccess, _ := rws.tryGetRedisLock(labelTreeLockKey); lockRedisSuccess {
		defer func() {
			// 防止key删不掉
			for i := 0; i < 5; i++ {
				if _, err := rws.redisClient.Del(context.Background(), labelTreeLockKey).Result(); err != nil && !errors.Is(err, redis.Nil) {
					log.Printf("删除key失败, key:%s , error:%#v", labelTreeLockKey, err)
				} else {
					break
				}
			}
		}()

		log.Println("+++++++++++++++++++执行更新++++++++++++++++++++")
		return rws.strategyService.UpdateSpecificLabelTreeById(
			models.LabelTree{Id: labelTreeId, ParentId: newParentLabelTreeId},
			service.LabelParentId,
		)
	}

	return nil
}

/*
  - kernelStrategyQuery 单例模式+分布式锁实现的查询逻辑
    场景读多写少

    前提条件：
    1. 现在我们的场景是单redis实例、单mysql结点！！！
    2. 由于redis是单线程的，那么我们读写就是分离的，我们可以不用使用读锁来保证数据在读的时候有写操作更新数据
    思路：
    1. 直接使用写锁来保证缓存写操作成功
    2. 优先使用本地锁，单实例并发是多实例并发的子集，所以优先本地加锁防止大量协程同时打入到redis中，造成无意义的分布式加锁
    3. 缓存key为空就查表插入数据，查表逻辑分为3层，优先本地加锁，然后redis加锁，最后查表
*/
func (rws *ReadWriteLockStrategy) kernelStrategyQuery(labelTreeId int) (string, error) {
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
	if success := rws.tryGetReadLocalLock(); success {
		defer rws.Unlock()
		//log.Println("成功获取到本地锁")

		//  在初始流量高的时候 下面这两步是最容易获取到缓存数据的
		val, err := rws.redisClient.Get(context.Background(), labelTreeKey).Result()
		if errors.Is(err, redis.Nil) {
			log.Println("第二次尝试获取缓存，仍然没有获取到缓存")
			val, err = rws.remoteConcurrentJudge(labelTreeKey, labelTreeId)
		}

		return val, err
	}

	// 2.2 本地获锁失败 直接返回 表示并发量特别大rws.localLockNum++
	// 因为已经阻塞了30s了再尝试获取一次缓存
	// 在初始流量高的时候 下面这两步是最容易获取到缓存数据的
	if val, err := rws.redisClient.Get(context.Background(), labelTreeKey).Result(); err == nil {
		atomic.AddInt32(&rws.localLockNum, 1)
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

	if lockRedisSuccess, _ := rws.tryGetRedisLock(labelTreeLockKey); lockRedisSuccess {
		defer func() {
			// 防止key删不掉
			for i := 0; i < 5; i++ {
				if _, err := rws.redisClient.Del(context.Background(), labelTreeLockKey).Result(); err != nil && !errors.Is(err, redis.Nil) {
					log.Printf("删除key失败, key:%s , error:%#v", labelTreeLockKey, err)
				} else {
					break
				}
			}
		}()

		// 4.1 获得分布式锁成功 第三次判断是否有缓存
		val, err := rws.redisClient.Get(context.Background(), labelTreeKey).Result()
		if errors.Is(err, redis.Nil) {
			log.Println("成功获取获取分布式锁 准备更新缓存")
			val, err = rws.strategyService.PushSpecificLabelTreeInfoById(labelTreeId)
			rws.final++
		}
		return val, err
	}

	return "", errors.New("多实例竞争锁超时")
}

/*
  - tryGetReadLocalLock 获取本地锁操作
    追加计时器 如果阻塞时间超过30s那么就直接拒绝❌
    如果获得锁成功就返回true  否则继续尝试
*/
func (rws *ReadWriteLockStrategy) tryGetReadLocalLock() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	for {
		if success := rws.TryLock(); success {
			return true
		}

		// 上下文超时判断
		select {
		case <-ctx.Done():
			return false
		default:
			time.Sleep(time.Millisecond * 100)
			continue
		}
	}
}

/*
  - tryGetRedisLock 尝试获取redis的分布式锁
    使用带超时的上下文 当分布式锁获取时间超过30s 则获锁失败
*/
func (rws *ReadWriteLockStrategy) tryGetRedisLock(labelTreeLockKey string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for {
		success, err := rws.redisClient.SetNX(ctx, labelTreeLockKey, 1, time.Minute*2).Result()
		if err != nil {
			log.Printf("redis 报错了: %#v\n", err)
			return false, err
		}

		// 获取锁成功
		if success {
			return true, nil
		}

		// 上下文超时判断
		select {
		case <-ctx.Done():
			return false, nil
		default:
			time.Sleep(time.Millisecond * 100)
			continue
		}
	}
}

// updateLabelInformation 更新label的信息
func (rws *ReadWriteLockStrategy) updateLabelInformation(labelTreeId int) {

}
