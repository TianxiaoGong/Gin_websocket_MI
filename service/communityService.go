package service

import (
	"Gin_WebSocket_IM/dao"
	"Gin_WebSocket_IM/models"
	"Gin_WebSocket_IM/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

// CreateCommunity
// @Summary 新建群
// @Tags 群模块
// @param ownerId query string false "用户id"
// @param name query string false "群名称"
// @Success 200 {string} json{"code","msg","data“}
// @Router /contact/createCommunity [post]
func CreateCommunity(c *gin.Context) {
	cd := dao.NewCommunityDao()
	ownerId, _ := strconv.Atoi(c.Request.FormValue("ownerId"))
	name := c.Request.FormValue("name")
	community := models.Community{}
	community.OwnerId = uint(ownerId)
	community.Name = name
	code, msg := cd.CreateCommunity(community)
	if code == 0 {
		utils.Success(c, msg, nil)
	} else {
		utils.Failed(c, msg)
	}
}

// LoadCommunity
// @Summary 加载群列表
// @Tags 群模块
// @param ownerId query string false "用户id"
// @param name query string false "群名称"
// @Success 200 {string} json{"code","msg","data“}
// @Router /contact/LoadCommunity [post]
func LoadCommunity(c *gin.Context) {
	cd := dao.NewCommunityDao()
	ownerId, _ := strconv.Atoi(c.Request.FormValue("ownerId"))
	data, msg := cd.LoadCommunity(uint(ownerId))
	if len(data) == 0 {
		utils.RespOKList(c.Writer, data, msg)
	} else {
		utils.RespFail(c.Writer, msg)
	}
}
