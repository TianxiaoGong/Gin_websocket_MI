package dao

import (
	"Gin_WebSocket_IM/models"
	"Gin_WebSocket_IM/utils"
	"fmt"
)

type ContactDao struct {
}

func NewContactDao() *ContactDao {
	return &ContactDao{}
}

func (cd *ContactDao) SearchFriend(ownerId uint) []models.UserBasic {
	contacts := make([]models.Contact, 0)
	objIds := make([]uint64, 0)
	utils.DB.Where("owner_id = ? and type = 1", ownerId).Find(&contacts)
	for _, v := range contacts {
		fmt.Println(">>>>>>>>>>>>>>>>>>", v)
		objIds = append(objIds, uint64(v.TargetId))
	}
	users := make([]models.UserBasic, 0)
	utils.DB.Where("id in ?", objIds).Find(&users)
	return users
}

// AddFriend 添加好友
func (cd *ContactDao) AddFriend(ownerId uint, targetId uint) (int, string) {
	user := models.UserBasic{}
	ud := NewUserDao()
	if targetId != 0 {
		user = ud.FindUserByID(targetId)
		if user.Salt != "" {
			c := models.Contact{}
			utils.DB.Where("owner_id = ? and target_id = ? and type=1", ownerId, targetId).Find(&c)
			if c.ID != 0 {
				return -1, "该用户已添加"
			}

			tx := utils.DB.Begin()
			//事务一旦开始，不论什么异常最终都会Rollback
			defer func() {
				if r := recover(); r != nil {
					tx.Rollback()
				}
			}()
			contact := models.Contact{}
			contact_ := models.Contact{}
			contact.OwnerId = ownerId
			contact.TargetId = targetId
			contact.Type = 1
			if err := utils.DB.Create(&contact).Error; err != nil {
				tx.Rollback()
				return 1, "添加好友失败"
			}
			contact_.OwnerId = targetId
			contact_.TargetId = ownerId
			contact_.Type = 1
			if err := utils.DB.Create(&contact_).Error; err != nil {
				tx.Rollback()
				return 1, "添加好友失败"
			}
			tx.Commit()
			return 0, "添加好友成功"
		}
		return 1, "没有找到此用户"
	}
	return 1, "好友ID不能为空"
}
