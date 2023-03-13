package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/haisin-official/haisin/app/validators"
	"github.com/haisin-official/haisin/config"
	"github.com/haisin-official/haisin/database"
	"github.com/haisin-official/haisin/routes"
	_ "github.com/lib/pq"
)

func main() {
	// Connect to postgresql
	dbType := config.GetEnv("DB_TYPE")
	dbDSN := config.GetEnv("DB_DSN")
	database.InitClient(dbType, dbDSN)
	defer database.CloseClient()

	// Start Gin Framework
	r := gin.Default()

	// Add routes for Gin Framework
	routes.Router(r)

	// Register Validators for Gin Framework
	validators.Register()

	port := config.GetEnv("HTTP_PORT")
	r.Run(":" + port)

	fmt.Println("hais.in version:", config.GetEnv("VERSION"))
	fmt.Println("Listening:", port)
}
