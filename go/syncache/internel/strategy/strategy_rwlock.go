package strategy

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"sync"
	"syncache/internel/client"
	"syncache/internel/models"
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
	Context Context
	BaseStrategy
	mysqlClient *gorm.DB
	redisClient *redis.Client

	labelTreeDao *models.LabelTreeMapper
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
		rws.mysqlClient = client.MysqlInstance.Get(rws.Context.config)
		rws.redisClient = client.RedisInstance.Get(rws.Context.config)

		rws.labelTreeDao = models.NewLabelTreeDao(rws.mysqlClient)
	})
}

// run strategy 的run方法接口
func (rws *ReadWriteLockStrategy) run() {
	log.Println("单例模式（读写锁实现）协程号：", utils.GetGoroutineID())
	log.Println("redis key upload start ")

	labelTrees := rws.labelTreeDao.GetAllLabelTree()
	if len(labelTrees) != 0 {
		redisNodeIdToParents := "%s:%d"
		redisNodeIdToParentsAll := "%s:all"
		mapIdToParents, mapIdToLabelTree := rws.mergeLabelTree(labelTrees)

		// 遍历插入数据
		for nodeId, parentIds := range mapIdToParents {
			redisLabelTreeDataTemp := template.NewRedisMapLabelTreeDataTemp(mapIdToLabelTree[nodeId].Name, parentIds)
			jsonData, err := json.Marshal(redisLabelTreeDataTemp)
			if err != nil {
				log.Println(err)
			}

			if errRedis := rws.redisClient.SetEx(context.Background(), fmt.Sprintf(redisNodeIdToParents, labelTreeKey, nodeId), string(jsonData), time.Minute*5).Err(); errRedis != nil {
				// 如果err 不为空 那么就要重试
				log.Println("failed!!")
			}
		}

		jsonMapIdToParents, err := json.Marshal(mapIdToParents)
		if err != nil {
			log.Println(err)
		}

		// 全量插入
		if errRedis := rws.redisClient.SetEx(context.Background(), fmt.Sprintf(redisNodeIdToParentsAll, labelTreeKey), string(jsonMapIdToParents), time.Minute*5).Err(); errRedis != nil {
			// 如果err 不为空 那么就要重试
			log.Println("failed!!")
		}
	}

	log.Println("redis key 上传成功🏅")
}
