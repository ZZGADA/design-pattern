package models

import (
	"gorm.io/gorm"
	"sync"
	"syncache/conf"
	"syncache/internel/client"
)

type LabelTreeMapper struct {
	sync.Once
	client *gorm.DB
}

type LabelTree struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	ParentId int    `json:"parent_id"`
	Level    int    `json:"level"`
}

var labelTreeDao *LabelTreeMapper

func init() {
	labelTreeDao = &LabelTreeMapper{}
}

// init 实例化对象 获取mysql客户端
func (lt *LabelTreeMapper) init() *LabelTreeMapper {
	lt.Do(func() {
		lt.client = client.MysqlInstance.Get(conf.Dft.Get())
	})
	return lt
}

// NewLabelTreeDao 初始化
func NewLabelTreeDao() *LabelTreeMapper {
	return labelTreeDao.init()
}

// TableName 表名
func (LabelTree) TableName() string {
	return "label_tree"
}

// GetAllLabelTree 获取全量的体系树
func (lt *LabelTreeMapper) GetAllLabelTree() []LabelTree {
	const rootId int = 1
	var labelTrees []LabelTree

	lt.client.Where("id != ?", rootId).Find(&labelTrees)
	return labelTrees
}

// GetById GetSpecificLabelTreeParentById 根据id查询指定一个label tree的父级信息
func (lt *LabelTreeMapper) GetById(labelTreeId int) LabelTree {
	labelTree := LabelTree{}

	lt.client.Where("id = ?", labelTreeId).Find(&labelTree)
	return labelTree
}

// UpdateLabelNameById 更新label体系的名称
func (lt *LabelTreeMapper) UpdateLabelNameById(labelTree LabelTree) error {
	result := lt.client.Model(&LabelTree{}).Where("id = ?", labelTree.Id).Updates(LabelTree{Name: labelTree.Name})
	return result.Error
}

// UpdateLabelParentId 更新label的父级id
func (lt *LabelTreeMapper) UpdateLabelParentId(labelTree LabelTree) error {
	result := lt.client.Model(&LabelTree{}).Where("id = ?", labelTree.Id).Updates(LabelTree{ParentId: labelTree.ParentId})
	return result.Error
}
