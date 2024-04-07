package main

import (
	"blue/controller"
	"blue/dao/mysql"
	"blue/dao/redis"
	"blue/logger"
	"blue/pkg/snowflake"
	"blue/router"
	"blue/setting"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	//if len(os.Args) < 2 {
	//	fmt.Println("need config file.eg:blue config.yaml")
	//}
	//
	//// 1. 加载配置
	//if err := setting.Init(os.Args[1]); err != nil {
	//	fmt.Printf("init setting failed,err: %v \n", err)
	//}

	// 1. 加载配置
	if err := setting.Init(); err != nil {
		fmt.Printf("init setting failed,err: %v \n", err)
	}

	// 2. 初始化日志
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed,err: %v \n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")

	// 3. 初始化Mysql连接
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed,err: %v \n", err)
	}

	// 4. 初始化Redis连接
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed,err: %v \n", err)
	}
	defer redis.Close()
	//初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed,err: %v /n", err)
		return
	}

	//初始化雪花算法
	if err := snowflake.Init("2023-12-27", 1); err != nil {
		fmt.Printf("init snowflake failed,err:%v \n", err)
		return
	}
	//注册路由
	r := router.SetupRouter(setting.Conf.Mode)

	//启动HTTP服务
	r.Run(":80")
}
