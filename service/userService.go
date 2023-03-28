package service

import (
	"Gin_WebSocket_IM/dao"
	"Gin_WebSocket_IM/models"
	"Gin_WebSocket_IM/utils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
)

// GetUserList
// @Summary 获取所有用户信息
// @Tags 用户模块
// @Success 200 {string} json{"code","msg","data“}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	//data := make([]*models.UserBasic, 10)
	ud := dao.NewUserDao()
	data := ud.GetUserList()
	utils.Success(c, "获取用户列表成功", data)
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","msg","data“}
// @Router /user/createUser [post]
func CreateUser(c *gin.Context) {
	ud := dao.NewUserDao()
	user := models.UserBasic{}
	user.Name = c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	repassword := c.Request.FormValue("Identity")
	fmt.Println(user.Name, "<<<<", password, "<<<", repassword)
	//user.Name = c.Query("name")
	//password := c.Query("password")
	//repassword := c.Query("repassword")
	user.Salt = fmt.Sprintf("%06d", rand.Int31())
	findUser := ud.FindUserByName(user.Name)
	if user.Name == "" || password == "" {
		utils.Failed(c, "用户名或密码不能为空")
		return
	}
	if findUser.Name != "" {
		utils.Failed(c, "用户名已注册")
		return
	}
	if password != repassword {
		utils.Failed(c, "两次密码不一致")
		return
	}
	user.PassWord = utils.MakePassword(password, user.Salt)
	ud.CreateUser(user)
	utils.Success(c, "新增用户成功", user)
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} json{"code","msg","data“}
// @Router /user/deleteUser [post]
func DeleteUser(c *gin.Context) {
	ud := dao.NewUserDao()
	user := models.UserBasic{}
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		utils.Failed(c, "缺失删除用户id信息")
		return
	}
	user.ID = uint(ID)
	ud.DeleteUser(user)
	utils.Success(c, "删除用户成功", user)
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code","msg","data“}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	ud := dao.NewUserDao()
	user := models.UserBasic{}
	ID, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(ID)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	if _, err := govalidator.ValidateStruct(user); err != nil {
		fmt.Println(err)
		utils.Failed(c, "修改参数不匹配")
	} else {
		ud.UpdateUser(user)
		utils.Success(c, "修改用户成功", user)
	}
}

// LoginByNameAndPwd
// @Summary 用户登录
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","msg","data“}
// @Router /user/loginByNameAndPwd [post]
func LoginByNameAndPwd(c *gin.Context) {
	ud := dao.NewUserDao()
	//user := models.UserBasic{}
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	//repassword := c.Query("repassword")
	//if password != repassword {
	//	utils.Failed(c, "两次密码不一致")
	//	return
	//}
	findUser := ud.FindUserByName(name)
	if findUser.Name == "" {
		utils.Failed(c, "该用户不存在")
		return
	}
	if !utils.ValidPassword(password, findUser.Salt, findUser.PassWord) {
		utils.Failed(c, "密码不正确")
		return
	}
	data := ud.LoginByNameAndPwd(name, findUser.PassWord)
	utils.Success(c, "登录成功", data)
}

// SearchFriends
// @Summary 查询好友列表
// @Tags 用户模块
// @param userId query string false "用户id"
// @Success 200 {string} json{"code","msg","data“}
// @Router /user/searchFriends [post]
func SearchFriends(c *gin.Context) {
	cd := dao.NewContactDao()
	ID, _ := strconv.Atoi(c.Request.FormValue("userId"))
	users := cd.SearchFriend(uint(ID))
	utils.RespOKList(c.Writer, users, len(users))
}
