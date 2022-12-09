package main

import (
	"fmt"
	"gin_demo/logger"
	"gin_demo/setting"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func main() {

	// 1. 加载配置viper
	//if len(os.Args) < 2 {
	//	panic("程序执行时必须通过命令行指定配置文件")
	//}
	//fmt.Println(os.Args)
	//err := setting.Init(os.Args[1])
	err := setting.InitConfig("E:\\helloGolang\\src\\gin_demo\\config\\config.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Println(setting.Conf.Name + " " + setting.Conf.Version)

	// 2.初始化日志模块zap
	err = logger.InitLogger()
	if err != nil {
		panic(err)
	}
	defer zap.L().Sync()
	//defer logger.Logger.Sync()

	logger.Logger.Info("app start")

	r := gin.Default()
	// 测试接口
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	r.Run(setting.Conf.Address)

}
