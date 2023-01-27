package v1

import (
	"github.com/gin-gonic/gin"
	"gonovel/model"
	"gonovel/utils/errMsg"
	"net/http"
	"strconv"
)

func AddComment(c *gin.Context) {
	var comment model.Comments
	_ = c.ShouldBindJSON(&comment)
	code := model.AddComment(comment)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errMsg.GetErrMsg(code),
	})
}

func DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userId, _ := strconv.Atoi(c.Param("user_id"))
	code := model.DeleteComment(id, userId)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errMsg.GetErrMsg(code),
	})
}

func GetComments(c *gin.Context) {
	article := c.Query("article_id")
	title := c.Query("title")

	comments, total := model.GetComments(article, title)
	if total > 0 {
		code = errMsg.SUCCESS
	} else {
		code = errMsg.CommentNotExist
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"message":  errMsg.GetErrMsg(code),
		"comments": comments,
		"total":    total,
	})
}

func AddAgree(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userId, _ := strconv.Atoi(c.Param("user_id"))
	var comment model.Comments
	_ = c.ShouldBindJSON(&comment)
	code := model.AddAgree(id, userId, comment.Agrees)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errMsg.GetErrMsg(code),
	})
}
