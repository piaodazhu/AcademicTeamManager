package global

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Conf        Config
	MysqlClient *gorm.DB
	RedisClient *redis.Client
)
