package main

import (
	"fmt"
	"gin_demo/dao/mysql"
	"gin_demo/logger"
	"gin_demo/router"
	"gin_demo/setting"
	"go.uber.org/zap"
)

func main() {

	// 1. 加载配置viper
	// 命令行启动: go run main.go "E:\\helloGolang\\src\\gin_demo\\config\\config.yaml"
	// if len(os.Args) < 2 {
	//	panic("程序执行时必须通过命令行指定配置文件")
	//}
	//fmt.Println(os.Args)
	//err := setting.InitConfig(os.Args[1])
	err := setting.InitConfig("E:\\helloGolang\\src\\gin_demo\\config\\config.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Println(setting.Conf.Name + " " + setting.Conf.Version)

	// 2.初始化日志模块zap
	err = logger.Init()
	if err != nil {
		panic(err)
	}
	defer zap.L().Sync()
	//defer logger.Logger.Sync()
	zap.L().Info("app start")
	//logger.Logger.Info("app start")

	// 3. 数据库初始化
	mysql.Init()

	// 4. 初始化路由
	r := router.Setup()

	// 5. 程序启动
	r.Run(setting.Conf.Address)

}
