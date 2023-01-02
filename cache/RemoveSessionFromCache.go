package cache

import (
	"fmt"
)

func RemoveSessionFromCache(userId string, sessionId string) {
	query := MakeAccessTokenKey(userId, sessionId, "*")
	matches := DB.Keys(Ctx, query).Val()

	for _, match := range matches {
		_, err := DB.Del(Ctx, match).Result()
		if err != nil {
			fmt.Println(err)
		}
	}
}
