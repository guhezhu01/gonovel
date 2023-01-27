package v1

import (
	"github.com/gin-gonic/gin"
	"gonovel/model"
	"gonovel/utils/errMsg"
	"net/http"
	"strconv"
)

func GetUserInformation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("userId"))
	data, code := model.GetUserInformation(id)
	c.JSON(http.StatusOK, gin.H{
		"status":      code,
		"information": data,
		"message":     errMsg.GetErrMsg(code),
	})

}

func CreateUserInformation(c *gin.Context) {
	var data model.UserInformation
	_ = c.ShouldBindJSON(&data)
	code := model.CreateUserInformation(data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errMsg.GetErrMsg(code),
	})

}
func EditUserInformation(c *gin.Context) {
	var data model.UserInformation
	_ = c.ShouldBindJSON(&data)

	code := model.EditUserInformation(data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errMsg.GetErrMsg(code),
	})
}
