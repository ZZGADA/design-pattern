package strategy

import (
	"syncache/internel/models"
)

// Strategy 策略模式
type Strategy interface {
	run()
	init()
}

type BaseInterfaceStrategy interface {
	mergeLabelTree(labelTrees []models.LabelTree) (map[int]string, map[int]models.LabelTree)                     //	组装label树
	dfsFindParentLabelTree(mapIdToParents map[int]string, mapIdToLabelTree map[int]models.LabelTree, nodeId int) //  dfs递归+记忆化搜索
}

const (
	labelTreeKey = "label:tree"
)
