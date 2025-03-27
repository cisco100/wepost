package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cisco100/wepost/internal/store"
	"github.com/go-redis/redis/v8"
)

type UserCache struct {
	rdb *redis.Client
}

func (uc *UserCache) Get(ctx context.Context, userID string) (*store.User, error) {
	cacheKey := fmt.Sprintf("user-%s", userID)
	ctx, cancel := context.WithTimeout(ctx, store.QueryTimeoutDuration)
	defer cancel()
	data, err := uc.rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, nil

	}

	var user store.User
	if data != "" {
		err := json.Unmarshal([]byte(data), &user)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil

}

func (uc *UserCache) Set(ctx context.Context, user *store.User) error {
	cacheKey := fmt.Sprintf("user-%s", user.ID)

	json, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return uc.rdb.SetEX(ctx, cacheKey, json, time.Minute).Err()

}
