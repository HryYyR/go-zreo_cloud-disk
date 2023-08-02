package svc

import (
	"cloud_disk/go-zreo_cloud-disk/core/internal/config"
	"cloud_disk/go-zreo_cloud-disk/core/internal/middleware"
	"cloud_disk/go-zreo_cloud-disk/core/models"

	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/rest"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Xorm   *xorm.Engine
	Redis  *redis.Client
	Logger rest.Middleware
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Xorm:   models.InitXorm(c.Mysql.DataSource),
		Redis:  models.InitRedis(c),
		Logger: middleware.NewLoggerMiddleware().Handle,
		Auth:   middleware.NewAuthMiddleware().Handle, //token中间件
	}
}
