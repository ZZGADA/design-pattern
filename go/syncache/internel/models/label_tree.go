package models

import (
	"gorm.io/gorm"
)

type LabelTreeMapper struct {
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

func NewLabelTreeDao(db *gorm.DB) *LabelTreeMapper {
	labelTreeDao.client = db
	return labelTreeDao
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
