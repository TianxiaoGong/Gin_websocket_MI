package dao

import (
	"Gin_WebSocket_IM/models"
	"Gin_WebSocket_IM/utils"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type UserDao struct {
	*utils.DataBase
}

func NewUserDao() *UserDao {
	return &UserDao{utils.DB}
}

// GetUserList 获取用户列表
func (ud *UserDao) GetUserList() []*models.UserBasic {
	data := make([]*models.UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func (ud *UserDao) LoginByNameAndPwd(name, password string) models.UserBasic {
	user := models.UserBasic{}
	utils.DB.Where("name = ? and pass_word = ?", name, password).First(&user)
	//token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.MD5Encode(str)
	utils.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp)
	return user
}

func (ud *UserDao) FindUserByID(id uint) models.UserBasic {
	user := models.UserBasic{}
	utils.DB.Where("id = ?", id).First(&user)
	return user
}

func (ud *UserDao) FindUserByName(name string) models.UserBasic {
	user := models.UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}

func (ud *UserDao) FindUserByPhone(phone string) models.UserBasic {
	user := models.UserBasic{}
	utils.DB.Where("phone = ?", phone).First(&user)
	return user
}

func (ud *UserDao) FindUserByEmail(email string) models.UserBasic {
	user := models.UserBasic{}
	utils.DB.Where("email = ?", email).First(&user)
	return user
}

func (ud *UserDao) CreateUser(user models.UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func (ud *UserDao) DeleteUser(user models.UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

func (ud *UserDao) UpdateUser(user models.UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(models.UserBasic{
		Name:     user.Name,
		PassWord: user.PassWord,
		Phone:    user.Phone,
		Email:    user.Email,
	})
}
