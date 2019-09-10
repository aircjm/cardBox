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
	auth := auth{}
	c.BindJSON(&auth)
	data := make(map[string]interface{})
	isExist := authDao.CheckAuth(auth.Username, auth.Password)
	code := 0
	if isExist.ID > 0 {
		token, err := util.GenerateToken(auth.Username, auth.Password)
		if err != nil {
			code = responseStatus.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token
			data["uuid"] = isExist.ID
			data["name"] = auth.Username
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
