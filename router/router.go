package router

import (
	"gin_demo/controller"
	"gin_demo/controller/books"
	"gin_demo/controller/products"
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

	// 基础路由拆分
	r.GET("/hello", controller.HelloHandler)
	r.GET("/ping", controller.PingHandler)
	r.GET("/login", controller.LoginHandler)

	// 路由拆分为不同模块
	books.Init(r)

	products.Init(r)

	return r
}
