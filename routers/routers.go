package routers

import (
	"github.com/aircjm/gocard/common"
	"github.com/aircjm/gocard/common/responseStatus"
	"github.com/aircjm/gocard/controller/authController"
	"github.com/aircjm/gocard/controller/boardController"
	"github.com/aircjm/gocard/controller/cardController"
	"github.com/aircjm/gocard/middleware/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 初始化路由规则
func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/")

	api.POST("/login", authController.GetAuth)

	// 业务模块API-卡片服务
	cardGroup := api.Group("/card", jwt.JWT())

	{
		cardGroup.GET("/saveRecentCard", cardController.SaveRecentCard)
		cardGroup.POST("/getCardList", cardController.GetCardList)
		cardGroup.POST("/convertToAnki", cardController.ConvertToAnki)
		cardGroup.POST("/updateCardStatus", cardController.UpdateCardStatus)
		cardGroup.GET("/saveAllCards", cardController.SaveAllCards)

	}
	// 业务模块API-卡片服务
	boardGroup := api.Group("/board", jwt.JWT())

	{
		boardGroup.GET("/getBoardList", boardController.GetBoardList)
	}

	// Index
	//router.StaticFile("/", "./dist/index.html")
	//router.StaticFile("/index.html", "./dist/index.html")
	//router.StaticFile("/index.htm", "./dist/index.html")

	// Assets that should be cached
	//router.StaticFile("/favicon.ico", "./dist/favicon.ico")
	//router.Static("/js", "./dist/js")
	//router.Static("/css", "./dist/css")
	//router.Static("/img", "./dist/img")
	//router.Static("/image", "./dist/image")
	//router.Static("/fonts", "./dist/fonts")

	return router
}

// Ping API
func Ping(c *gin.Context) {
	cG := common.Gin{C: c}
	cG.Response(http.StatusOK, responseStatus.SUCCESS, "goCard服务正常")
}
