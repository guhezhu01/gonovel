package v1

import (
	"github.com/gin-gonic/gin"
	"gonovel/model"
	"gonovel/utils/errMsg"
	"net/http"
)

func GetArticle(c *gin.Context) {
	var article model.Article
	_ = c.ShouldBindJSON(&article)
	article, code := model.GetArticle(article.Title)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errMsg.GetErrMsg(code),
		"article": article,
	})
}

func GetTypeArticles(c *gin.Context) {
	type_ := c.Query("type")
	articles, total, code := model.GetTypeArticles(type_)
	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"message":  errMsg.GetErrMsg(code),
		"total":    total,
		"articles": articles,
	})
}

func GetRandArticle(c *gin.Context) {
	data, total, code := model.GetRandArticle()
	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"message":  errMsg.GetErrMsg(code),
		"total":    total,
		"articles": data,
	})

}
func DownloadArticle(c *gin.Context) {
	id := c.Query("id")
	title := c.Query("title")
	contents, code := model.GetArticleContents(id, title)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errMsg.GetErrMsg(code),
		"data":    contents,
	})
}

func DeleteArticle(c *gin.Context) {
	var article model.Article
	_ = c.ShouldBindJSON(&article)
	code := model.DeleteArticle(article.Id, article.Type)
	code01 := model.DeleteArticleContents(article.Id, article.Type)
	if code == errMsg.SUCCESS && code01 == errMsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  errMsg.SUCCESS,
			"message": errMsg.GetErrMsg(errMsg.SUCCESS),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  errMsg.ERROR,
			"message": errMsg.GetErrMsg(errMsg.ERROR),
		})
	}

}
