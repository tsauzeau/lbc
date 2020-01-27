package db

import (
	"github.com/go-redis/redis"
)

var client *redis.Client

//Init ...
func Init(addr string) {
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

//GetClient ...
func GetClient() *redis.Client {
	return client
}
