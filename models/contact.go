package models

import (
	"gorm.io/gorm"
)

// Contact 人员关系
type Contact struct {
	gorm.Model
	OwnerId  uint   //谁的关系
	TargetId uint   //对应的谁
	Type     int    //对应的类型 1好友 2群
	Desc     string //描述
}

func (table *Contact) TableName() string {
	return "contact"
}
