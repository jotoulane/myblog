package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"web01/settings"
)

var client *redis.Client

func Init() (err error) {

	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%d",
			settings.Conf.RedisConfig.Host,
			settings.Conf.RedisConfig.Port,
		),
		Password: settings.Conf.RedisConfig.Password, // 密码
		DB:       settings.Conf.RedisConfig.DB,       // 数据库
		PoolSize: settings.Conf.PoolSize,             // 连接池大小
	})
	_, err = client.Ping().Result()
	return
}
