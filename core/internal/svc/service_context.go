package svc

import (
	"cloud_disk/core/internal/config"
	"cloud_disk/core/internal/middleware"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"os"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config      config.Config
	Engine      *xorm.Engine
	RDB         *redis.Client
	EmailSender *Sender
	CosSecret   *TencentSecret
	Auth        rest.Middleware
}

type Sender struct {
	Email   string
	AuthPwd string
}

type TencentSecret struct {
	SecretID  string
	SecretKey string
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
		Engine: Init(c.Mysql.DataSource),
		RDB:    InitRedis(c),
		EmailSender: &Sender{
			Email:   os.Getenv("EmailSender"),
			AuthPwd: os.Getenv("EmailAuthPwd"),
		},
		CosSecret: &TencentSecret{
			SecretID:  os.Getenv("TencentSecretID"),
			SecretKey: os.Getenv("TencentSecretKey"),
		},
		Auth: middleware.NewAuthMiddleware().Handle,
	}
}
