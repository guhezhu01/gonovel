package v1

import (
	"github.com/gin-gonic/gin"
	"gonovel/middleware"
	"gonovel/model"
	"gonovel/utils/errMsg"
	"net/http"
)

func Login(c *gin.Context) {
	var userData model.User
	_ = c.ShouldBindJSON(&userData)
	var token string
	var code int
	userData, code = model.CheckLogin(userData.Username, userData.Password)
	if code == errMsg.SUCCESS {
		token, _ = middleware.SetToken(userData.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":          code,
		"message":         errMsg.GetErrMsg(code),
		"token":           token,
		"userInformation": userData,
	})
}
