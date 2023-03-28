package service

import (
	"Gin_WebSocket_IM/dao"
	"Gin_WebSocket_IM/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"text/template"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	ind, err := template.ParseFiles("index.html", "view/chat/head.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "index")
}

// ToRegister
// @Tags 首页
// @Success 200 {string} welcome
// @Router /toRegister [get]
func ToRegister(c *gin.Context) {
	ind, err := template.ParseFiles("view/user/register.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "register")
}

// ToChat
// @Tags 首页
// @Success 200 {string} welcome
// @Router /toChat [get]
func ToChat(c *gin.Context) {
	ind, err := template.ParseFiles("view/chat/index.html",
		"view/chat/head.html",
		"view/chat/tabmenu.html",
		"view/chat/concat.html",
		"view/chat/group.html",
		"view/chat/profile.html",
		"view/chat/main.html",
		"view/chat/foot.html")
	if err != nil {
		panic(err)
	}
	userId, _ := strconv.Atoi(c.Query("userId"))
	token := c.Query("token")
	user := models.UserBasic{}
	user.ID = uint(userId)
	user.Identity = token
	fmt.Println("Tochat>>>", user)
	ind.Execute(c.Writer, user)
}

func Chat(c *gin.Context) {
	md := dao.NewMsgDao()
	md.Chat(c.Writer, c.Request)
}
