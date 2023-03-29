package app

import (
	"Gin_WebSocket_IM/docs"
	"Gin_WebSocket_IM/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	//swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//静态资源
	r.Static("/asset", "asset/")
	r.LoadHTMLGlob("view/**/*")
	//首页
	r.GET("/", service.GetIndex)
	r.GET("/index", service.GetIndex)
	r.GET("/toRegister", service.ToRegister)
	r.GET("/toChat", service.ToChat)
	r.GET("/chat", service.Chat)
	//用户模块
	r.GET("/user/getUserList", service.GetUserList)
	r.POST("/user/createUser", service.CreateUser)
	r.POST("/user/deleteUser", service.DeleteUser)
	r.POST("/user/updateUser", service.UpdateUser)
	r.POST("/user/loginByNameAndPwd", service.LoginByNameAndPwd)
	r.POST("/user/searchFriends", service.SearchFriends)
	r.POST("/user/addFriend", service.AddFriends)
	//发送消息
	r.GET("/msg/sendMsg", service.SendMsg)
	r.GET("/msg/sendUserMsg", service.SendUserMsg)
	//上传文件
	r.POST("/attach/upload", service.Upload)
	//创建群
	r.POST("/contact/createCommunity", service.CreateCommunity)
	//群列表
	r.POST("/contact/loadCommunity", service.LoadCommunity)

	//读取redis缓存信息
	r.POST("/user/redisMsg", service.RedisMsg)
	return r
}
