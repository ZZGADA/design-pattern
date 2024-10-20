package strategy

import (
	"strconv"
	"syncache/internel/models"
)

// BaseStrategy strategy基类
type BaseStrategy struct {
}

/*
*
  - mergeLabelTree 拼接出体系树
    期望形式：
    1. key 为每个结点，value为父级的所有结点
    2. value是一个拼接的字符串，形如：a,b,c,d
    3. 字符串的拼接顺序从最高级节点到最当前结点
    4. 如果父级结点为1表示为层级最高的父级结点 （level 1）
    return:
    1. 返回树状结构的map
*/
func (s *BaseStrategy) mergeLabelTree(labelTrees []models.LabelTree) (map[int]string, map[int]models.LabelTree) {
	mapIdToLabelTree := map[int]models.LabelTree{}
	mapIdToParents := map[int]string{}
	length := len(labelTrees)

	// 拼接出字符串
	for i := 0; i < length; i++ {
		mapIdToLabelTree[labelTrees[i].Id] = labelTrees[i]
	}

	// dfs 递推
	for nodeIdKey := range mapIdToLabelTree {
		s.dfsFindParentLabelTree(mapIdToParents, mapIdToLabelTree, nodeIdKey)
	}
	return mapIdToParents, mapIdToLabelTree
}

// dfsFindParentLabelTree dfs+记忆化搜索 建树
func (s *BaseStrategy) dfsFindParentLabelTree(mapIdToParents map[int]string, mapIdToLabelTree map[int]models.LabelTree, nodeId int) {

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
		s.dfsFindParentLabelTree(mapIdToParents, mapIdToLabelTree, parentId)
	}
	mapIdToParents[nodeId] = mapIdToParents[parentId] + "," + strconv.Itoa(nodeId)
	return
}
