package impl

import (
	"errors"
	"log"
	"strconv"
	"sync"
	"syncache/internel/models"
	"syncache/internel/service"
)

type LabelTreeService struct {
	sync.Once
	labelTreeDao *models.LabelTreeMapper
}

var labelTreeService *LabelTreeService

func init() {
	labelTreeService = &LabelTreeService{}
}

func (lbs *LabelTreeService) init() *LabelTreeService {
	lbs.Do(func() {
		lbs.labelTreeDao = models.NewLabelTreeDao()
	})
	return lbs
}

func NewLabelTreeService() service.LabelTreeService {
	return labelTreeService.init()
}

/*
MergeLabelTreeAll  拼接出体系树

期望形式：
1. key 为每个结点，value为父级的所有结点
2. value是一个拼接的字符串，形如：a,b,c,d
3. 字符串的拼接顺序从最高级节点到最当前结点
4. 如果父级结点为1表示为层级最高的父级结点 （level 1）

return:
1. 返回树状结构的map
*/
func (lbs *LabelTreeService) MergeLabelTreeAll() (map[int]string, map[int]models.LabelTree) {
	labelTrees := lbs.labelTreeDao.GetAllLabelTree()
	mapIdToLabelTree := make(map[int]models.LabelTree)
	mapIdToParents := make(map[int]string)
	length := len(labelTrees)

	// 将slice => 变为 map
	for i := 0; i < length; i++ {
		mapIdToLabelTree[labelTrees[i].Id] = labelTrees[i]
	}

	// dfs 递推
	for nodeIdKey := range mapIdToLabelTree {
		lbs.DfsFindParentLabelTree(mapIdToParents, mapIdToLabelTree, nodeIdKey)
	}
	return mapIdToParents, mapIdToLabelTree
}

/*
MergeLabelTreeOne 针对一个node结点拼出一个体系树

模式：
使用了适配器模式

算法：
使用bfs搜索当前结点下面的所有子结点

问题：
与上面的差异只在于穿了一个参数
只有一个结点 那么我们需要 更新这一个结点、其子结点
子结点是未知的 我们无法不通过递归的形式找到当前结点和自结点的关系
所以直接调用 MergeLabelTreeAll
然后筛选出相关联的当前结点、子结点
*/
func (lbs *LabelTreeService) MergeLabelTreeOne(nodeKeyId int) (map[int]string, map[int]models.LabelTree) {
	mapIdToParentsAll, mapIdToLabelTree := lbs.MergeLabelTreeAll()
	mapIdToParentChildren := make(map[int]string)
	needUpdateId := make([]int, 0)
	queue := make([]int, 0)
	needUpdateId = append(needUpdateId, nodeKeyId)
	queue = append(queue, nodeKeyId)

	for len(queue) > 0 {
		queueSize := len(queue)
		for queueSize > 0 {
			queueHead := queue[0]
			queue = queue[1:]

			//  bfs 全局遍历
			for node, nodeInfo := range mapIdToLabelTree {
				if nodeInfo.ParentId == queueHead {
					queue = append(queue, node)
					needUpdateId = append(needUpdateId, node)
				}
			}

			queueSize--
		}
	}

	for i := 0; i < len(needUpdateId); i++ {
		mapIdToParentChildren[needUpdateId[i]] = mapIdToParentsAll[needUpdateId[i]]
	}
	return mapIdToParentChildren, mapIdToLabelTree
}

/*
DfsFindParentLabelTree

dfs+记忆化搜索 建树
当前结点想上级结点递归，如果上级结点存在就直接返回追加，否则继续递归上级结点 直至返回
*/
func (lbs *LabelTreeService) DfsFindParentLabelTree(mapIdToParents map[int]string, mapIdToLabelTree map[int]models.LabelTree, nodeId int) {

	parentId := mapIdToLabelTree[nodeId].ParentId
	if parentId == 1 {
		// 终止条件
		// 父级id为1 表示到达最高的层级	level1
		mapIdToParents[nodeId] = strconv.Itoa(parentId) + "," + strconv.Itoa(nodeId)
		return
	}

	// 1. key不存在 表示父级的层级还没有建立好
	// 	那么就有限建立父级的状态 然后拼接
	// 2. 否则key存在表示父级的树状结构已经处理好了
	if _, exist := mapIdToParents[parentId]; !exist {
		lbs.DfsFindParentLabelTree(mapIdToParents, mapIdToLabelTree, parentId)
	}
	mapIdToParents[nodeId] = mapIdToParents[parentId] + "," + strconv.Itoa(nodeId)
	return
}

// UpdateSpecificLabelTreeById 更新label tree的信息
func (lbs *LabelTreeService) UpdateSpecificLabelTreeById(labelTree models.LabelTree, attributes service.LabelAttributes) error {
	switch attributes {
	case service.LabelName:
		if err := lbs.labelTreeDao.UpdateLabelNameById(labelTree); err != nil {
			log.Println(err)
			return err
		}
	case service.LabelParentId:
		if err := lbs.labelTreeDao.UpdateLabelParentId(labelTree); err != nil {
			log.Println(err)
			return err
		}
	default:
		return errors.New("无效更新属性")
	}
	return nil
}
