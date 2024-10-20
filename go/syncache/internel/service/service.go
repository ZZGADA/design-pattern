package service

import "syncache/internel/models"

type StrategyService interface {
	//GetAllLabelTreeAllParent(labelTreeId int)       	// 获取全量的label tree 的所有父级信息
	GetSpecificLabelTreeInfoById(labelTreeId int) string // 根据id查询指定一个label tree的父级信息
	UpdateSpecificLabelTreeParentById(labelTreeId int)   // 更新一个体系标签
	//AddLabelTreeOne()                               	// 增加一个label 体系
	//AddsLabelTreeMulti()                            	// 增加多个体系标签
}

type LabelTreeService interface {
	MergeLabelTree() (map[int]string, map[int]models.LabelTree)                                                  //	组装label树
	DfsFindParentLabelTree(mapIdToParents map[int]string, mapIdToLabelTree map[int]models.LabelTree, nodeId int) //  dfs递归+记忆化搜索
}
