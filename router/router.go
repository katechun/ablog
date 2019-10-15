package router

import (
	"ablog/controllers"
	"ablog/helpers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine{
	//设置非debug模式
	gin.SetMode(gin.DebugMode)

	//设置路由
	router := gin.Default()

	//配置html模板
	helpers.SetTemplate(router)

	//设置静态文件路径
	router.Static("/static",helpers.GetOneValue("app","staticdir"))
	//如果没有正确路由 返回404信息
	router.NoRoute(controllers.Handle404)

	router.GET("/index",controllers.Index)

	router.GET("/test",controllers.Test)

	return router
}

