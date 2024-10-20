package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
	"syncache/conf"
	"syncache/internel/client"
	"syncache/internel/models"
	"syncache/internel/service"
	"syncache/internel/template"
	"time"
)

type StrategyService struct {
	sync.Once
	sync.Mutex

	labelTreeDao     *models.LabelTreeMapper
	redisClient      *redis.Client
	labelTreeService service.LabelTreeService
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
		s.labelTreeService = NewLabelTreeService()
	})
	return s
}

// NewStrategyService  初始化 多态
func NewStrategyService() service.StrategyService {
	return strategyService.init()
}

// PushSpecificLabelTreeInfoById 更新对应label和其子对象的的缓存
func (s *StrategyService) PushSpecificLabelTreeInfoById(labelTreeId int) (string, error) {
	log.Println("redis key upload start ")
	redisNodeIdToParents := "%s:%d"
	mapIdToParentsNeedPush, mapIdToLabelTree := s.labelTreeService.MergeLabelTreeOne(labelTreeId)

	// 遍历插入数据
	for nodeId, parentIds := range mapIdToParentsNeedPush {
		redisLabelTreeDataTemp := template.NewRedisMapLabelTreeDataTemp(mapIdToLabelTree[nodeId].Name, parentIds)
		jsonData, err := json.Marshal(redisLabelTreeDataTemp)
		if err != nil {
			log.Println(err)
			return "", err
		}

		if err := s.redisClient.SetEx(context.Background(), fmt.Sprintf(redisNodeIdToParents, template.RedisKeyLabelTree, nodeId), string(jsonData), time.Second*60).Err(); err != nil {
			// 如果err 不为空 那么就要重试
			log.Println(err)
			return "", err
		}
	}
	log.Println("redis 当前 key的所有变化 上传成功🏅")
	return mapIdToParentsNeedPush[labelTreeId], nil
}

// PushAllLabelTreeAllParent 全量插入
func (s *StrategyService) PushAllLabelTreeAllParent() map[int]string {
	redisNodeIdToParentsAll := "%s:all"
	mapIdToParentsAll, _ := s.labelTreeService.MergeLabelTreeAll()

	// 转json
	jsonMapIdToParents, err := json.Marshal(mapIdToParentsAll)
	if err != nil {
		log.Println(err)
	}

	// 全量插入
	if errRedis := s.redisClient.SetEx(context.Background(), fmt.Sprintf(redisNodeIdToParentsAll, template.RedisKeyLabelTree), string(jsonMapIdToParents), time.Minute*5).Err(); errRedis != nil {
		// 如果err 不为空 那么就要重试
		log.Println("failed!!")
	}

	log.Println("redis 全量 key 上传成功🏅")
	return mapIdToParentsAll
}

// pushCacheToRedis 将缓存推入redis
func (s *StrategyService) pushCacheToRedis() {

}

func (s *StrategyService) UpdateSpecificLabelTreeParentById(labelTreeId int) {

}
