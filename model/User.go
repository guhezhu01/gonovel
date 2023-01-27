package model

import (
	"gonovel/utils/errMsg"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

func CheckUser(username string) (code int) {
	var users User
	db.Select("id").Where("username", username).First(&users)
	if users.ID > 0 {
		return errMsg.UsernameUsed
	}
	return errMsg.SUCCESS
}

func CreateUser(userData *User) (code int) {
	err = db.Create(&userData).Error
	if err != nil {
		return errMsg.ERROR //500
	}
	return errMsg.SUCCESS
}

func GetUsers() ([]User, int) {
	var users []User
	var total int

	err = db.Model(&users).Find(&users).Error
	total = len(users)

	if err != nil {
		return users, 0
	}
	return users, total
}
func GetUser(username string) (User, int) {
	var user User
	err := db.Model(user).Where("username =?", username).Find(&user).Error
	if err != nil || user.Username == "" {
		return user, errMsg.UserNotExist
	}
	return user, errMsg.SUCCESS
}
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}

func EditUser(id int, userData *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = userData.Username
	maps["role"] = userData.Role
	err := db.Model(&user).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}

func CheckLogin(username string, password string) (User, int) {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.Username == "" {
		return user, errMsg.UserNotExist
	} else if user.Password != password {
		return user, errMsg.PasswordWrong
	}
	return user, errMsg.SUCCESS
}

func UpdateUserPassword(username, password string, id uint) int {
	err := db.Model(&User{}).Where("username =? and id =?", username, id).Update("password", password).Error
	if err != nil {
		return errMsg.UpdatePasswordWrong
	}
	return errMsg.SUCCESS
}
