package cache

import (
	"api/models"
	"fmt"
)

func RemoveUserFromCache(user *models.User) {
	matches := DB.Keys(Ctx, user.ID.String()+":*").Val()

	for _, match := range matches {
		_, err := DB.Del(Ctx, match).Result()
		if err != nil {
			fmt.Println(err)
		}
	}
}
