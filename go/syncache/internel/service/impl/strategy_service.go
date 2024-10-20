package impl

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
	"syncache/conf"
	"syncache/internel/client"
	"syncache/internel/models"
	"syncache/internel/service"
	"syncache/internel/template"
)

type StrategyService struct {
	sync.Once

	labelTreeDao *models.LabelTreeMapper
	redisClient  *redis.Client
}

var strategyService *StrategyService

// 初始化
func init() {
	strategyService = &StrategyService{}
}

// init service 实例化
func (s *StrategyService) init() service.StrategyService {
	// 单例模式
	s.Do(func() {
		s.labelTreeDao = models.NewLabelTreeDao()
		s.redisClient = client.RedisInstance.Get(conf.Dft.Get())
	})
	return s
}

// NewStrategyService  初始化 多态
func NewStrategyService() service.StrategyService {
	return strategyService.init()
}

// GetSpecificLabelTreeInfoById 获取一个体系标签的信息
func (s *StrategyService) GetSpecificLabelTreeInfoById(labelTreeId int) string {
	labelTreeKey := fmt.Sprintf("%s:%d", template.RedisKeyLabelTree, labelTreeId)

	// 1. 优先获取缓存
	val, err := s.redisClient.Get(context.Background(), labelTreeKey).Result()
	if errors.Is(err, redis.Nil) {
		// key 不存在 需要从DB中查value

		return ""
	}

	// 1.2. 查缓存报错
	if err != nil {
		// redis 报错了
		log.Printf("redis获取的key报错，error：%#v\n", err)
		return "业务繁忙，稍等一下～"
	}

	// 2. 缓存存在就直接返回
	return val
}

func (s *StrategyService) UpdateSpecificLabelTreeParentById(labelTreeId int) {

}
