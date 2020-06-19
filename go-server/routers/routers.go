package routers

import (
	"github.com/aircjm/gocard/common"
	"github.com/aircjm/gocard/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 初始化路由规则
func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/")

	// TODO 缺少网关
	//cardGroup.Use(middleware.Jwt())

	// 公共模块API
	commonGroup := api.Group("/common")

	{
		commonGroup.GET("/login", controller.Login)
	}
	// 业务模块API-卡片服务
	cardGroup := api.Group("/card")

	{
		cardGroup.GET("/saveRecentCard", controller.SaveRecentCard)
		cardGroup.GET("/saveAllCards", controller.SaveAllCards)
		cardGroup.POST("/getCardList", controller.GetCardList)
		cardGroup.POST("/convertToAnki", controller.ConvertToAnki)
	}
	// 业务模块API-卡片服务
	boardGroup := api.Group("/board")

	{
		boardGroup.GET("/getBoardList", controller.GetBoardList)
		boardGroup.GET("/saveAllCards", controller.SaveAllCards)
	}

	// Index
	router.StaticFile("/", "./dist/index.html")
	router.StaticFile("/index.html", "./dist/index.html")
	router.StaticFile("/index.htm", "./dist/index.html")

	// Assets that should be cached
	router.StaticFile("/favicon.ico", "./dist/favicon.ico")
	router.Static("/js", "./dist/js")
	router.Static("/css", "./dist/css")
	router.Static("/img", "./dist/img")
	router.Static("/image", "./dist/image")
	router.Static("/fonts", "./dist/fonts")

	return router
}

// Ping API
func Ping(c *gin.Context) {
	cG := common.Gin{C: c}
	cG.Response(http.StatusOK, common.Success, "goCard服务正常")
}
