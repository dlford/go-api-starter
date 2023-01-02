package cache

import (
	"api/constants"
	"context"

	"github.com/go-redis/redis/v8"
)

var DB *redis.Client
var Ctx = context.Background()

func ConnectCache() {
	DB = redis.NewClient(&redis.Options{
		Addr:     constants.REDIS_HOST,
		Password: constants.REDIS_PASSWORD,
		DB:       constants.REDIS_DB,
	})
}
