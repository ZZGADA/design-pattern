package strategy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"sync"
	"syncache/conf"
	"syncache/define"
	"syncache/internel/client"
	"syncache/internel/models"
	"syncache/internel/service"
	"syncache/internel/service/impl"
	"syncache/internel/utils"
	"time"
)

// DoubleDeleteStrategy 延时双删策略
type DoubleDeleteStrategy struct {
	sync.Once
	sync.Mutex
	Context Context
	BaseStrategy
	redisClient *redis.Client

	labelTreeDao     *models.LabelTreeMapper
	labelTreeService service.LabelTreeService
	strategyService  service.StrategyService
}

// NewDoubleDeleteStrategy 初始化对象
func NewDoubleDeleteStrategy(context Context) *DoubleDeleteStrategy {
	strategy := &DoubleDeleteStrategy{Context: context}
	strategy.init()
	return strategy
}

func (dds *DoubleDeleteStrategy) init() {
	dds.Do(func() {
		log.Println("DoubleDeleteStrategy 懒加载执行依赖注入～ ，单例加载")
		dds.redisClient = client.RedisInstance.Get(conf.Dft.Get())
		dds.labelTreeDao = models.NewLabelTreeDao()
		dds.labelTreeService = impl.NewLabelTreeService()
		dds.strategyService = impl.NewStrategyService()
	})
}

/*
- run 延时双删策略
*/
func (dds *DoubleDeleteStrategy) run() {
	log.Println("延时双删策略 开始 协程号：", utils.GetGoroutineID())

	maxIteratorNum := 33
	stopStep := 10
	wg := new(sync.WaitGroup)
	for i := 0; i < maxIteratorNum; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			labelTree, _ := dds.kernelStrategyQuery(context.Background(), 10)
			log.Info(labelTree)
		}()

		if i == stopStep {
			if err := dds.kernelStrategyUpdate(context.Background(), 10); err != nil {
				log.Errorf("kernelStrategyUpdate error %#v", err)
			} else {
				log.Info("update success ,label tree name is ZZGDEA")
			}
		}
	}
	wg.Wait()
	log.Println("延时双删策略 结束")
}

/*
  - kernelStrategyQuery 延时双删 - 查询策略
    场景： 读多写少，只使用一个mysql实例情况下最佳，或者使用云厂商提供的成熟的mysql集群

    思路：
    1. 读操作时，缓存不存在则从mysql读数据，然后更新redis
    2. 写操作时，先更新mysql，然后删除缓存，然后间隔1段时间后再删除缓存数据
    3. 删除两次的原因是，第一次删除缓存后，由于读请求是从从库读数据，但是存在主从库的延时，于是导致从从库中读到老数据，将老数据推入redis缓存。
    于是延时一段时间（主从库同步时间）再将缓存删除，防止从库读取的老数据缓存一致存在在redis中。
    4. 对于向阿里的polarDB这种数据库，主从延时极小，使用延时双删既可以简化代码，又不容易出错
    5. 额外注意一下⚠️：不能在事务中将mysql数据推入redis。因为MVCC的事务隔离机制，导致事务的select语句不会拿到最新的数据，而是一个快照。一旦将这个快照数据推入redis将直接导致脏数据的出现。
*/
func (dds *DoubleDeleteStrategy) kernelStrategyQuery(context context.Context, labelTreeId int) (labelTree models.LabelTree, err error) {
	// 1. 查缓存
	labelTreeRedisKey := fmt.Sprintf(define.DoubleDeleteKey, labelTreeId)
	ex := dds.redisClient.Get(context, labelTreeRedisKey)
	labelTreeVal := ex.Val()
	if ex.Err() != nil && !errors.Is(ex.Err(), redis.Nil) {
		log.Errorf("获取redis缓存失败 ， err：%#v", ex.Err())
	}

	// 2. 缓存不存在
	// 多携程情况下 存在多携程查询redis key 不存在 导致大量请求进入mysql
	if len(labelTreeVal) == 0 {
		// value 不存在 就查数据入缓存
		labelTree = dds.labelTreeDao.GetById(labelTreeId)
		if labelTreeJson, errM := json.Marshal(labelTree); errM == nil {
			dds.redisClient.SetEx(context, labelTreeRedisKey, string(labelTreeJson), time.Second*60)
		} else {
			err = errM
		}

		log.Info("从数据库中拿数据 ")
	} else {
		// 3. 缓存存在
		if err = json.Unmarshal([]byte(labelTreeVal), &labelTree); err != nil {
			log.Errorf("label tree缓存解析成结构体失败 %#v", err)
		}
		log.Info("从缓存中拿数据 ")
	}

	return
}

/*
- kernelStrategyUpdate 延时双删 - 更新策略
思路：

1. 更新数据库后删除redis缓存，然后间隔一段时间（主从库）的同步时间再删除缓存。
2. 延时删除的意义在于 防止读操作将从库的老数据读入缓存中。
*/
func (dds *DoubleDeleteStrategy) kernelStrategyUpdate(context context.Context, labelTreeId int) (err error) {
	// 1. 更新数据库
	labelTree := dds.labelTreeDao.GetById(labelTreeId)
	labelTree.Name = "ZZGEDA"
	if err = dds.labelTreeService.UpdateSpecificLabelTreeById(labelTree, service.LabelName); err != nil {
		log.Errorf("label tree 更新失败 ，labelTree ：%#v  error: %#v", labelTree, err)
		return
	}

	// 2. 删除缓存
	labelTreeRedisKey := fmt.Sprintf(define.DoubleDeleteKey, labelTreeId)
	delRes := dds.redisClient.Del(context, labelTreeRedisKey)
	if delRes.Err() != nil && !errors.Is(delRes.Err(), redis.Nil) {
		log.Errorf("redis 缓存删除失败，labelTree id :%#v,err：%#v ", labelTreeId, delRes.Err())
		return
	}

	// 3. 延时双删
	time.Sleep(define.Delay1Second)
	delResDouble := dds.redisClient.Del(context, labelTreeRedisKey)
	if delResDouble.Err() != nil && !errors.Is(delRes.Err(), redis.Nil) {
		log.Errorf("redis 缓存删除失败，labelTree id :%#v,err：%#v ", labelTreeId, delResDouble.Err())
		return
	}

	return
}
