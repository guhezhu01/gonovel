package v1

import (
	"github.com/gin-gonic/gin"
	"gonovel/model"
	"gonovel/utils/errMsg"
	"net/http"
	"strconv"
)

var code int

// 获取所有用户
func GetUsers(c *gin.Context) {

	data, total := model.GetUsers()
	code := errMsg.SUCCESS
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errMsg.GetErrMsg(code),
		},
	)
}

// 查询单个用户
func GetUser(c *gin.Context) {
	username := c.Query("username")
	userData, code := model.GetUser(username)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    userData,
		"message": errMsg.GetErrMsg(code),
	})
}

// 添加新用户
func AddUser(c *gin.Context) {
	var userData model.User
	_ = c.ShouldBindJSON(&userData)
	code = model.CheckUser(userData.Username)
	if code == errMsg.SUCCESS {
		code = model.CreateUser(&userData)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    userData,
		"message": errMsg.GetErrMsg(code),
	})
}

// 删除新用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errMsg.GetErrMsg(code),
	})
}

// 修改用户
func EditUser(c *gin.Context) {
	var userData model.User
	_ = c.ShouldBindJSON(&userData)
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.CheckUser(userData.Username)
	if code == errMsg.SUCCESS {
		code = model.EditUser(id, &userData)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errMsg.GetErrMsg(code)})
}

func UpdateUserPassword(c *gin.Context) {
	var userData model.User
	_ = c.ShouldBindJSON(&userData)
	code1 := model.CheckUser(userData.Username)
	if code1 != 200 {
		code := model.UpdateUserPassword(userData.Username, userData.Password, userData.ID)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errMsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errMsg.GetErrMsg(errMsg.UserNotExist),
		})
	}

}
