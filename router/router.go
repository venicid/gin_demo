package router

import (
	"gin_demo/controller"
	"gin_demo/middleware"
	"go.uber.org/zap"

	"gin_demo/setting"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	gin.SetMode(setting.Conf.Mode)

	r := gin.New()

	// 注册两个自定义的中间件
	r.Use(middleware.Cors())
	r.Use(middleware.GinLogger(zap.L()))
	r.Use(middleware.GinRecovery(zap.L(), false))

	r.GET("/hello", controller.HelloHandler)

	return r
}
