package svc

import (
	"cloud_disk/core/internal/config"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: Init(c.Mysql.DataSource),
		RDB:    InitRedis(c),
	}
}
