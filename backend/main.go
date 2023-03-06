package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/haisin-official/haisin/config"
	"github.com/haisin-official/haisin/ent"
	"github.com/haisin-official/haisin/routes"
)

func main() {
	// Connect to postgresql
	dbType := config.GetEnv("DB_TYPE")
	dbDSN := config.GetEnv("DB_DSN")

	client, err := ent.Open(dbType, dbDSN)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	r := gin.Default()
	routes.Router(r)

	port := config.GetEnv("HTTP_PORT")
	r.Run(":" + port)

	fmt.Println("hais.in version:", config.GetEnv("VERSION"))
	fmt.Println("Listening:", port)
}
