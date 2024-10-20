package service

type Service interface {
	GetAllLabelTreeAllParent(labelTreeId int)       // 获取全量的label tree 的所有父级信息
	GetSpecificLabelTreeParentById(labelTreeId int) // 根据id查询指定一个label tree的父级信息
	AddLabelTreeOne()                               // 增加一个label 体系
	AddsLabelTreeMulti()                            // 增加多个体系标签
}
