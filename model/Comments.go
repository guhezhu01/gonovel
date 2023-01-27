package model

import (
	"fmt"
	"gonovel/utils/errMsg"
	"gorm.io/gorm"
)

// 评论
type Comments struct {
	gorm.Model
	Id        int    `json:"id"`
	UserId    uint   `json:"user_id"`
	ArticleId string `json:"article_id"`
	Title     string `json:"article_title"`
	Username  string `json:"username"`
	Content   string `json:"content"`
	Img       string `json:"img"`
	Agrees    int    `gorm:"default:0" json:"agrees"` //点赞数
	Target    string `json:"target"`                  //回复对象
	Pid       int    `json:"pid"`                     //父级评论id
}

func AddComment(comment Comments) int {
	err := db.Select("ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Id", "UserId",
		"ArticleId", "Title", "Username", "Content", "Agrees", "Target", "Pid").Create(&comment).Error
	if err != nil {
		fmt.Println("数据库添加失败", err)
		return errMsg.ERROR //500
	}

	return errMsg.SUCCESS
}

func DeleteComment(id int, userId int) int {
	var comment Comments
	err := db.Where("id = ? and user_id = ?", id, userId).Delete(&comment).Error
	if err != nil {
		return errMsg.ERROR
	}

	return errMsg.SUCCESS
}

func GetComments(articleId, title string) ([]Comments, int64) {
	var comments []Comments
	var total int64
	err := db.Raw("select * from comments,user_information where comments.user_id=user_information.user_id "+
		"and article_id=? and title =?", articleId, title).Scan(&comments).Error
	//err := db.Where("article_id=? and title =?", articleId, title).Find(&comments).Count(&total).Error
	if err != nil {
		return comments, 0
	}
	total = int64(len(comments))
	return comments, total
}

func AddAgree(id, userId, agrees int) int {
	err := db.Model(&Comments{}).Where("id=? and user_id =?", id, userId).Update("agrees", agrees).Error
	if err != nil {
		return errMsg.CommentAddWrong
	}
	return errMsg.SUCCESS
}
