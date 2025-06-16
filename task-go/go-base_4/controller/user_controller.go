package controller

import (
	"github.com/gin-gonic/gin"
	"task-go/task-go/go-base_4/models"
	"task-go/task-go/go-base_4/response"
)

func UserRegister(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		return
	}

	err = models.Register(user)

	if err != nil {
		response.FailWithMsg(c, "user already exists")
	} else {
		response.OkWithMsg(c, "register successfully")
	}
	return
}

func UserLogin(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		return
	}

	//生成JWT
	token, err := models.Login(c, user)
	if err != nil {
		response.FailWithMsg(c, "token generate failed")
		return
	}
	response.OkWithData(c, token)
}
