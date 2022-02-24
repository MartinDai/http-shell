package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"http-shell/controller"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 8080, "端口号")
	flag.Parse()

	// 注册路由
	router := SetupRouter()
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.POST("/sh", controller.ExecuteShShell)
	router.POST("/execShell", controller.ExecuteShell)

	return router
}
