package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
}

func (c *Cache) Set(key string, value string) {
	ctx := context.Background()
	c.client.Set(ctx, key, value, 0)
}

func (c *Cache) Get(key string) string {
	ctx := context.Background()
	value, _ := c.client.Get(ctx, key).Result()
	return value
}

func NewCache() *Cache {
	return &Cache{redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})}
}
