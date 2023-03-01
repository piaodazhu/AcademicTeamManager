package initializer

import (
	"atm/global"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func initRedis() {
	conf := global.Conf.RedisConf
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB: conf.DB,
	})
}
