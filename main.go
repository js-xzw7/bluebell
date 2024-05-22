package main

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/pkg/sonyflake"
	"bluebell/routers"
	"bluebell/settings"
	"fmt"
)

func main() {

	//加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()

	if err := redis.Init(*settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	if err := sonyflake.Init(1); err != nil {
		fmt.Printf("init sonyflake failed, err:%v\n", err)
		return
	}

	r := routers.SetupRouter()
	address := fmt.Sprintf(":%d", settings.Conf.Port)
	if err := r.Run(address); err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
