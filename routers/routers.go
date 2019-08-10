package routers

import (
	"github.com/aircjm/gocard/common"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 初始化路由规则
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", Ping)
	api := router.Group("/api/")

	// 公共模块API
	commonGroup := api.Group("/common")

	// 业务模块API-卡片服务
	cardGroup := api.Group("/card")

	log.Println(commonGroup)
	log.Println(cardGroup)

	// 网关注释
	//cardGroup.Use(middleware.Jwt())
	return router
}

// Ping API
func Ping(c *gin.Context) {
	cG := common.Gin{C: c}
	cG.Response(http.StatusOK, common.Success, "goCard服务正常")
}
