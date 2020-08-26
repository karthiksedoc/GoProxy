package redisInterface

import (
	"fmt"
	"log"
	"github.com/go-redis/redis"
)

func dom_search(name string) string, error {
	client := redis.NewClient(&redis.Options{
//Need to parameterize redis server credentials
		Addr: "redis.server.com:6379",
		Password:"",
		DB:0,
	})
	dom_stat, err := client.Get(name).Result()
	if err != nil {
		return "", err
	}
	return dom_stat, nil
}
