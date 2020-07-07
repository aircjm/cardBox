package controller

import (
	"github.com/aircjm/cardBox/common"
	"github.com/aircjm/cardBox/service"
	"github.com/gin-gonic/gin"
)

func GetBoardList(c *gin.Context) {
	cG := common.Gin{C: c}
	cG.Response(200, 0, service.GetBoardList())
}
