package dao

import (
	"Gin_WebSocket_IM/models"
	"Gin_WebSocket_IM/utils"
	"fmt"
)

type CommunityDao struct {
}

func NewCommunityDao() *CommunityDao {
	return &CommunityDao{}
}

func (cd *CommunityDao) CreateCommunity(community models.Community) (int, string) {
	tx := utils.DB.Begin()
	//事务一旦开始，不论什么异常最终都会 Rollback
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if community.OwnerId == 0 {
		return 1, "请先登录"
	}
	if len(community.Name) == 0 {
		return 1, "群名称不能为空"
	}
	if err := utils.DB.Create(&community).Error; err != nil {
		fmt.Println("CreateCommunity err:", err)
		return 1, "建群失败"
	}
	tx.Commit()
	return 0, "建群成功"
}

func (cd *CommunityDao) LoadCommunity(ownerId uint) ([]*models.Community, string) {
	data := make([]*models.Community, 10)
	utils.DB.Where("owner_id = ?", ownerId).Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data, "查询成功"
}
