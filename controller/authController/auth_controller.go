package authController

import (
	"github.com/aircjm/gocard/common/responseStatus"
	"github.com/aircjm/gocard/dao/authDao"
	"github.com/aircjm/gocard/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string
	Password string
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	data := make(map[string]interface{})
	isExist := authDao.CheckAuth(username, password)
	code := 0
	if isExist {
		token, err := util.GenerateToken(username, password)
		if err != nil {
			code = responseStatus.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token

			code = responseStatus.SUCCESS
		}

	} else {
		code = responseStatus.ERROR_AUTH_TOKEN
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  responseStatus.GetStatusMsg(code),
		"data": data,
	})
}
