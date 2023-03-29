package models

import (
	"gorm.io/gorm"
)

// Message 消息
type Message struct {
	gorm.Model
	FromId     int64  //发送者
	TargetId   int64  //接收者
	Type       int    //发送对象类型（1私聊 2群聊 3广播）
	Media      int    //消息类型（1文字 2表情包 3图片 4音频）
	Content    string //消息内容
	Pic        string
	Url        string
	Desc       string //描述
	Amount     int    //数字统计
	CreateTime uint64 //创建时间
	ReadTime   uint64 //读取时间
}

func (table *Message) TableName() string {
	return "message"
}
