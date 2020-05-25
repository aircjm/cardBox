package main

import (
	"github.com/aircjm/gocard/routers"
	"log"
)

func main() {

	// 开启定时器
	//go cronInit()

	// 初始化路由
	router := routers.InitRouter()
	router.Run(":" + "8085")

	// 项目启动日志
	log.Println("hello goCard")
}
