package boardController

import (
	"github.com/aircjm/gocard/common"
	"github.com/aircjm/gocard/service"
	"github.com/gin-gonic/gin"
)

func GetBoardList(c *gin.Context) {
	cG := common.Gin{C: c}
	cG.Response(200, 0, service.GetBoardList())
}
