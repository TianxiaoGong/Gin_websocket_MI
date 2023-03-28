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

func (cd *ContactDao) SearchFriend(userId uint) []models.UserBasic {
	contacts := make([]models.Contact, 0)
	objIds := make([]uint64, 0)
	utils.DB.Where("owner_id = ? and type = 1", userId).Find(&contacts)
	for _, v := range contacts {
		fmt.Println(">>>>>>>>>>>>>>>>>>", v)
		objIds = append(objIds, uint64(v.TargetId))
	}
	users := make([]models.UserBasic, 0)
	utils.DB.Where("id in ?", objIds).Find(&users)
	return users
}
