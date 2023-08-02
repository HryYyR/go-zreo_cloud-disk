package models

import (
	"cloud_disk/go-zreo_cloud-disk/core/internal/config"
	"context"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

func InitXorm(dataSoure string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", dataSoure)
	if err != nil {
		logx.Error(context.Background(), "xorm new engine error:", err)
	}
	return engine
}

func InitRedis(c config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func InitLogger() {
	cfg := new(logx.LogConf)
	cfg.Mode = "file"
	cfg.Path = "/core/logs"
	cfg.Encoding = "json"
	cfg.KeepDays = 30

	logx.MustSetup(*cfg)
	defer logx.Close()
}
