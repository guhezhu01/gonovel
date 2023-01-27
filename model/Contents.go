package model

import "gonovel/utils/errMsg"

type Contents struct {
	Title   string `gorm:"type:varchar(255);not null" json:"title"`
	Id      string `gorm:"type:varchar(20);not null" json:"id"`
	Type    string `gorm:"type:varchar(20);not null" json:"type"`
	Content string `gorm:"mediumtext" json:"content"`
}

func GetArticleContents(id string, title string) (string, int) {
	var contents Contents
	err := db.Where("id = ? and title = ?", id, title).Find(&contents).Error
	if err != nil {
		return "", errMsg.ArtNotExist
	}
	return contents.Content, errMsg.SUCCESS
}

func DeleteArticleContents(id string, type_ string) int {
	var contents Contents
	err := db.Unscoped().Where("id =? and type=?", id, type_).Delete(&contents).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}
