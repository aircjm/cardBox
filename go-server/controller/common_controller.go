package controller

import (
	"github.com/aircjm/gocard/common"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	// todo login方法缺少
	cG := common.Gin{C: c}

	cG.Response(200, 0, nil)

}
