package client

import (
	"fmt"
	_redis "github.com/redis/go-redis/v9"
	"sync"
	"syncache/conf"
)

var (
	RedisInstance *Redis
)

func init() {
	RedisInstance = &Redis{}
}

type Redis struct {
	sync.Once
	client *_redis.Client
}

type dataSourceRedis interface {
	Get(config conf.Config) *_redis.Client
}

func (redis *Redis) Get(config conf.Config) *_redis.Client {
	// 单例锁
	redis.Once.Do(func() {
		const url string = "%s:%s"
		redisConf := config.RedisConfig

		// 配置客户端
		redisClient := _redis.NewClient(&_redis.Options{
			Addr:           fmt.Sprintf(url, redisConf.Host, redisConf.Port),
			Password:       redisConf.Password,       // 没有密码，默认值
			DB:             redisConf.DataBase,       // 默认DB 0
			MaxActiveConns: redisConf.MaxActiveConns, // 连接池最大连接数
		})
		redis.client = redisClient
	})

	return redis.client
}
