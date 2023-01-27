package model

import "gonovel/utils/errMsg"

type UserCollection struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	ArticleId int    `json:"article_id"`
	Type      string `json:"type"`
}

func AddUserCollection(data UserCollection) int {
	err := db.Create(&data).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS

}

func GetUserCollections(userId int) ([]Article, int, int) {
	var userCollections []Article
	var total int
	err := db.Raw("select * from user_collection,article where article.id = user_collection.article_id "+
		"and article.type = user_collection.type and user_id =? ", userId).Scan(&userCollections).Error

	total = len(userCollections)
	if err != nil {
		return userCollections, 0, errMsg.ERROR
	}
	return userCollections, total, errMsg.SUCCESS
}

func DeleteUserCollection(userId, articleId int, type_ string) int {
	var userCollection UserCollection
	err := db.Where("user_id=? and article_id =? and type =? ", userId, articleId, type_).Delete(&userCollection).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS

}

func CheckUserCollection(userId int, articleId int, type_ string) (int, UserCollection) {
	var userCollection UserCollection
	err := db.Model(&userCollection).Where("user_id=? and article_id=? and type=?", userId, articleId, type_).Find(&userCollection).Error
	if err != nil {
		return errMsg.ERROR, userCollection
	}
	return errMsg.SUCCESS, userCollection
}
