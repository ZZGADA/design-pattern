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

// init proto 实例化
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
func (s *StrategyService) PushAllLabelTreeAllParent() (map[int]string, error) {
	redisNodeIdToParentsAll := "%s:all"
	mapIdToParentsAll, _ := s.labelTreeService.MergeLabelTreeAll()

	// 转json
	jsonMapIdToParents, err := json.Marshal(mapIdToParentsAll)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// 全量插入
	if err := s.redisClient.SetEx(context.Background(), fmt.Sprintf(redisNodeIdToParentsAll, template.RedisKeyLabelTree), string(jsonMapIdToParents), time.Minute*5).Err(); err != nil {
		// 如果err 不为空 那么就要重试
		log.Println(err)
		return nil, err
	}

	log.Println("redis 全量 key 上传成功🏅")
	return mapIdToParentsAll, nil
}

/*
UpdateSpecificLabelTreeById  根据label 的id更新label的信息 调用 label_tree proto

	更新数据库 然后将redis的key删除掉
	因为加锁了 ，两个删除key的操作是互补影响的所以直接并行执行 并执行重试操作
*/
func (s *StrategyService) UpdateSpecificLabelTreeById(labelTree models.LabelTree, attributes service.LabelAttributes) error {
	ch := make(chan error, 2)
	wg := new(sync.WaitGroup)
	defer func() {
		close(ch)
	}()

	mapIdToParentsNeedUpdateOld, _ := s.labelTreeService.MergeLabelTreeOne(labelTree.Id)
	if err := s.labelTreeService.UpdateSpecificLabelTreeById(labelTree, attributes); err != nil {
		return err
	}

	// 如果是更新名字的话 直接删除就好了
	if attributes == service.LabelName {
		labelTreeRedisKey := fmt.Sprintf("%s:%d", template.RedisKeyLabelTree, labelTree.Id)
		if _, err := s.redisClient.Del(context.Background(), labelTreeRedisKey).Result(); err != nil {
			return s.deleteMultiKeyReTry(labelTreeRedisKey)
		}
		return nil
	}

	// 否则是体系更新
	wg.Add(2)
	go func() {
		defer wg.Done()

		// 删除新的子集的缓存
		newLabelTreeKey := make([]string, 0)
		mapIdToParentsNeedUpdateNew, _ := s.labelTreeService.MergeLabelTreeOne(labelTree.Id)

		for nodeId, _ := range mapIdToParentsNeedUpdateNew {
			newLabelTreeKey = append(newLabelTreeKey, fmt.Sprintf("%s:%d", template.RedisKeyLabelTree, nodeId))
		}

		if _, err := s.redisClient.Del(context.Background(), newLabelTreeKey...).Result(); err != nil {
			ch <- s.deleteMultiKeyReTry(newLabelTreeKey...)
			return
		}
		ch <- nil
	}()

	go func() {
		defer wg.Done()

		// 删除旧缓存
		oldLabelTreeKey := make([]string, 0)
		for nodeId, _ := range mapIdToParentsNeedUpdateOld {
			oldLabelTreeKey = append(oldLabelTreeKey, fmt.Sprintf("%s:%d", template.RedisKeyLabelTree, nodeId))
		}

		if _, err := s.redisClient.Del(context.Background(), oldLabelTreeKey...).Result(); err != nil {
			ch <- s.deleteMultiKeyReTry(oldLabelTreeKey...)
			return
		}
		ch <- nil
	}()

	wg.Wait()

	for i := 0; i < 2; i++ {
		if err := <-ch; err != nil {
			return err
		}
	}

	return nil
}

// deleteMultiKeyReTry 删除key失败 重新尝试删key
func (s *StrategyService) deleteMultiKeyReTry(keyNeedReDelete ...string) error {
	var err error
	for i := 0; i < 5; i++ {
		if _, err = s.redisClient.Del(context.Background(), keyNeedReDelete...).Result(); err == nil {
			return nil
		}
		time.Sleep(time.Millisecond * 200)
	}
	return err
}
