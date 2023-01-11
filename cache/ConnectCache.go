package cache

import (
	"api/constants"
	"context"
	"crypto/tls"

	"github.com/go-redis/redis/v8"
)

var DB *redis.Client
var Ctx = context.Background()

func ConnectCache() {
	config := redis.Options{
		Addr:     constants.REDIS_HOST,
		Password: constants.REDIS_PASSWORD,
		DB:       constants.REDIS_DB,
	}

	if constants.REDIS_USER != "" {
		config.Username = constants.REDIS_USER
	}

	if constants.REDIS_SSL_ENABLED {
		config.TLSConfig = &tls.Config{}
	}

	DB = redis.NewClient(&config)
}
