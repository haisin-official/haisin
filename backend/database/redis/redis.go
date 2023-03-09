package redis

import (
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
)

func InitClient(url string) {
	opt, err := redis.ParseURL(url)
	if err != nil {
		log.Panicf("Fatal Error of redis 🚫\n %v", err)
	}
	client = redis.NewClient(opt)

	fmt.Println("Success connection to redis ✅")
}
func GetClient() *redis.Client {
	return client
}

func CloseClient() {
	client.Close()
}
