package models

import (
	"gorm.io/gorm"
)

// GroupBasic 群消息
type GroupBasic struct {
	gorm.Model
	Name    string
	OwnerId uint
	Icon    string //图标
	Type    int    //等级
	Desc    string //描述
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
