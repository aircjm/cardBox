package routers

import (
	"github.com/aircjm/cardBox/common"
	"github.com/aircjm/cardBox/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 初始化路由规则
func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/")
	router.GET("/ping", Ping)

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

	{
		// 业务模块API-卡片服务
		boardGroup := api.Group("/board")
		boardGroup.GET("/getBoardList", controller.GetBoardList)
		boardGroup.GET("/saveAllCards", controller.SaveAllCards)
	}

	util := api.Group("/util/")
	util.POST("/markdown2html", controller.Markdown2html)

	return router
}

// Ping API
func Ping(c *gin.Context) {
	cG := common.Gin{C: c}
	cG.Response(http.StatusOK, common.Success, "goCard服务正常")
}
