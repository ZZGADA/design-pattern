package strategy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"syncache/internel/client"
	"syncache/internel/models"
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
	Context     Context
	mysqlClient *gorm.DB
	redisClient *redis.Client

	labelTreeDao *models.LabelTreeMapper
}

// initStrategy 依赖注入初始化 构造器模式
func (rws *ReadWriteLockStrategy) initStrategy() {
	rws.Do(func() {
		log.Println("懒加载执行依赖注入～ ，单例加载")
		rws.mysqlClient = client.MysqlInstance.Get(rws.Context.config)
		rws.redisClient = client.RedisInstance.Get(rws.Context.config)

		rws.labelTreeDao = models.NewLabelTreeDao(rws.mysqlClient)
	})
}

func (rws *ReadWriteLockStrategy) run() {
	go func() {
		rws.initStrategy()
		log.Println("单例模式（读写锁实现）", getGoroutineID())
		ctx := context.Background()
		labelTrees := rws.labelTreeDao.GetAllLabelTree()
		jsonLabelTrees, err := json.Marshal(labelTrees)
		if err != nil {
			panic(err)
		}
		now := time.Now()
		rws.redisClient.SetEx(ctx, fmt.Sprintf("%s:%d", labelTreeKey, now.Unix()+rand.Int63()), string(jsonLabelTrees), time.Second*60)
		log.Println("redis key 上传成功🏅")
	}()
}

func getGoroutineID() uint64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := bytes.Fields(buf[:n])[1]
	id, err := strconv.ParseUint(string(idField), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
