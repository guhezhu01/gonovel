package v1

import (
	"github.com/gin-gonic/gin"
	"gonovel/model"
	"gonovel/utils/errMsg"
	"net/http"
	"strconv"
)

func GetUserCollections(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	userCollections, total, code := model.GetUserCollections(userId)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    userCollections,
		"total":   total,
		"message": errMsg.GetErrMsg(code),
	})
}

func AddUserCollection(c *gin.Context) {
	var userCollection model.UserCollection
	_ = c.ShouldBindJSON(&userCollection)

	_, data := model.CheckUserCollection(userCollection.UserId, userCollection.ArticleId, userCollection.Type)
	if data.Type != "" {
		c.JSON(http.StatusOK, gin.H{
			"status":  errMsg.CollectionAdded,
			"message": errMsg.GetErrMsg(errMsg.CollectionAdded),
		})
	} else {
		code := model.AddUserCollection(userCollection)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errMsg.GetErrMsg(code),
		})
	}

}

func DeleteUserCollection(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	articleId, _ := strconv.Atoi(c.Query("article_id"))
	type_ := c.Query("type")

	code := model.DeleteUserCollection(userId, articleId, type_)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errMsg.GetErrMsg(code),
	})
}
