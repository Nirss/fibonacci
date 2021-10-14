package redis_cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var (
	ErrUnexpectedError = errors.New("unexpected error, please repeat again later")
)

type Cache struct {
	client *redis.Client
}

func NewCache(port string) *Cache {
	address := fmt.Sprintf("localhost:%v", port)
	return &Cache{client: redis.NewClient(&redis.Options{Addr: address})}
}

func (r *Cache) GetValue(ctx context.Context, from int, to int) ([]int, error) {
	redisKey := fmt.Sprintf("from=%d to=%d", from, to)
	redisValue, err := r.client.Get(ctx, redisKey).Result()
	var result []int
	if err != nil {
		log.Println("get redis value error: ", err)
		return nil, ErrUnexpectedError
	}
	err = json.Unmarshal([]byte(redisValue), &result)
	if err != nil {
		log.Println("get redis value error: ", err)
		r.client.Del(ctx, redisKey)
		return nil, ErrUnexpectedError
	}
	return result, nil
}

func (r *Cache) SetValue(ctx context.Context, from int, to int, fibonacciValues []int) error {
	data, err := json.Marshal(fibonacciValues)
	if err != nil {
		log.Println("set redis value error: ", err)
		return ErrUnexpectedError
	}
	redisKey := fmt.Sprintf("from=%d to=%d", from, to)
	err = r.client.Set(ctx, redisKey, string(data), 0).Err()
	if err != nil {
		log.Println("set redis value error: ", err)
	}
	return nil
}
