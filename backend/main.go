package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/haisin-official/haisin/config"
	"github.com/haisin-official/haisin/database"
	"github.com/haisin-official/haisin/database/redis"
	"github.com/haisin-official/haisin/routes"
	_ "github.com/lib/pq"
)

func main() {
	// Connect to postgresql
	dbType := config.GetEnv("DB_TYPE")
	dbDSN := config.GetEnv("DB_DSN")
	database.InitClient(dbType, dbDSN)
	defer database.CloseClient()

	// Connect to redis
	redisUrl := config.GetEnv("REDIS_URL")
	redis.InitClient(redisUrl)
	defer redis.CloseClient()

	// Start Gin Framework
	r := gin.Default()

	// Start New Session Manager with redis
	redisClient := redis.GetClient()
	config.NewSessions(redisClient)

	// Add routes for Gin Framework
	routes.Router(r)

	port := config.GetEnv("HTTP_PORT")
	r.Run(":" + port)

	fmt.Println("hais.in version:", config.GetEnv("VERSION"))
	fmt.Println("Listening:", port)
}
