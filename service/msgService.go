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

// RedisMsg
// @Summary 读取Redis信息
// @Tags 信息模块
// @Router /user/redisMsg [get]
func RedisMsg(c *gin.Context) {
	md := dao.NewMsgDao()
	userIdA, _ := strconv.Atoi(c.Request.FormValue("userIdA"))
	userIdB, _ := strconv.Atoi(c.Request.FormValue("userIdB"))
	fmt.Println("userIDA::", userIdA)
	fmt.Println("userIDB::", userIdB)
	RdbMsg := md.RedisMsg(int64(userIdA), int64(userIdB))
	utils.RespOK(c.Writer, RdbMsg, "读取缓存信息成功")
}
