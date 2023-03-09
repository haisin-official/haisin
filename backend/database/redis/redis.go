package redis

import (
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	r *redis.Client
)

func InitClient(url string) {
	opt, err := redis.ParseURL(url)
	if err != nil {
		log.Panicf("Fatal Error of redis 🚫\n %v", err)
	}
	r = redis.NewClient(opt)

	fmt.Println("Success connection to redis ✅")
}
func GetClient() *redis.Client {
	return r
}

func CloseClient() {
	r.Close()
}
