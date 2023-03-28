package models

import (
	"gorm.io/gorm"
	"time"
)

// UserBasic 用户信息
type UserBasic struct {
	gorm.Model
	Name          string    `gorm:"name"`
	PassWord      string    `gorm:"pass_word"`
	Phone         string    `gorm:"phone" valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string    `gorm:"email" valid:"email"`
	Identity      string    `gorm:"identity"`
	ClientIP      string    `gorm:"client_ip"`
	ClientPort    string    `gorm:"client_port"`
	Salt          string    `gorm:"salt"`
	LoginTime     time.Time `gorm:"login_time"`
	HeartbeatTime time.Time `gorm:"heartbeat_time"`
	LoginOutTime  time.Time `gorm:"login_out_time"`
	IsLoginOut    bool      `gorm:"is_login_out"`
	DeviceInfo    string    `gorm:"device_info"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
