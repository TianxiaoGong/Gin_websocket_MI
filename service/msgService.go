package service

import (
	"Gin_WebSocket_IM/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
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
// @Summary 发送信息
// @Tags 信息模块
// @Router /msg/sendUserMsg [get]
func SendUserMsg(c *gin.Context) {
	md := dao.NewMsgDao()
	md.Chat(c.Writer, c.Request)
}
