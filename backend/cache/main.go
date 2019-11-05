package cache

import (
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client

// 用户sessionID索引前缀
var SessionIdxPrefix = "session:user:"

// Redis 在中间件中初始化redis链接
func Redis() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PW"),
		DB:       int(db),
		//PoolSize:80,
		MaxRetries:  2,
		IdleTimeout: time.Minute * 3,
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic(err)
	}

	RedisClient = client
}
