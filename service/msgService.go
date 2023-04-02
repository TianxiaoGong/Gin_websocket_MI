package service

import (
	"Gin_WebSocket_IM/dao"
	"Gin_WebSocket_IM/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

// 防止跨域站点伪造请求
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// SendMsg
// @Summary 发送信息
// @Tags 信息模块
// @Router /msg/sendMsg [get]
func SendMsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	md := dao.NewMsgDao()
	md.MsgHandler(ws, c)
}

// SendUserMsg
// @Summary 发送用户信息
// @Tags 信息模块
// @Router /msg/sendUserMsg [get]
func SendUserMsg(c *gin.Context) {
	md := dao.NewMsgDao()
	md.Chat(c.Writer, c.Request)
}

// GetRedisMsg
// @Summary 读取Redis信息
// @Tags 信息模块
// @param userIdA query string false "用户A_id"
// @param userIdB query string false "用户B_id"
// @Success 200 {string} json{"code","msg","data“}
// @Router /getRedisMsg [post]
func GetRedisMsg(c *gin.Context) {
	md := dao.NewMsgDao()
	userIdA, _ := strconv.Atoi(c.Request.FormValue("userIdA"))
	userIdB, _ := strconv.Atoi(c.Request.FormValue("userIdB"))
	fmt.Println("userIDA::", userIdA)
	fmt.Println("userIDB::", userIdB)
	RdbMsg := md.GetRedisMsg(int64(userIdA), int64(userIdB))
	utils.Success(c, "读取缓存成功", RdbMsg)
	//utils.RespOK(c.Writer, RdbMsg, "读取缓存信息成功")
}

// DeleteRedisMsg
// @Summary 读取Redis信息
// @Tags 信息模块
// @param userIdA query string false "用户A_id"
// @param userIdB query string false "用户B_id"
// @param nums query string false "删除记录数目"
// @Success 200 {string} json{"code","msg","data“}
// @Router /deleteRedisMsg [post]
func DeleteRedisMsg(c *gin.Context) {

}
