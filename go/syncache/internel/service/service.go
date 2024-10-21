package service

import (
	"syncache/internel/models"
)

type StrategyService interface {
	PushAllLabelTreeAllParent() (map[int]string, error)                                       // 获取全量的label tree 的所有父级信息
	PushSpecificLabelTreeInfoById(labelTreeId int) (string, error)                            // 根据id查询指定一个label tree的父级信息
	UpdateSpecificLabelTreeById(labelTree models.LabelTree, attributes LabelAttributes) error // 更新一个体系标签
	//AddLabelTreeOne()                               					// 增加一个label 体系
	//AddsLabelTreeMulti()                            					// 增加多个体系标签
}

type LabelTreeService interface {
	MergeLabelTreeAll() (map[int]string, map[int]models.LabelTree)                                               //	组装label树
	MergeLabelTreeOne(nodeKeyId int) (map[int]string, map[int]models.LabelTree)                                  // 查出单个结点的子结点 适配器模式
	DfsFindParentLabelTree(mapIdToParents map[int]string, mapIdToLabelTree map[int]models.LabelTree, nodeId int) // dfs递归+记忆化搜索
	UpdateSpecificLabelTreeById(labelTree models.LabelTree, attributes LabelAttributes) error                    // 更新labelTree
}
