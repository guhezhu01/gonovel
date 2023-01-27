package model

import (
	"gonovel/utils/errMsg"
)

type UserInformation struct {
	UserId        int    `json:"user_id"`
	Img           string `json:"img"`
	Desc          string `json:"desc"`
	Name          string `json:"name"`
	Communication string `json:"communication"`
}

func GetUserInformation(userId int) (UserInformation, int) {
	var data UserInformation
	err := db.Where("user_id=?", userId).Find(&data).Error
	if err != nil {
		return data, errMsg.UserNotExist
	}
	return data, errMsg.SUCCESS
}

func CreateUserInformation(data UserInformation) int {
	err := db.Create(&data).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}

func EditUserInformation(data UserInformation) int {
	err := db.Model(&data).Where("user_id=?", data.UserId).Updates(data).Error
	if err != nil {
		return errMsg.ERROR
	}

	return errMsg.SUCCESS
}
