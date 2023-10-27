package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/haisin-official/haisin/app/validators"
	"github.com/haisin-official/haisin/database"
	"github.com/haisin-official/haisin/routes"
	_ "github.com/lib/pq"
)

func main() {
	// Connect to postgresql
	dbType := os.Getenv("DB_TYPE")
	dbDSN := os.Getenv("DB_DSN")
	database.InitClient(dbType, dbDSN)
	defer database.CloseClient()

	// Start Gin Framework
	r := gin.Default()

	// Add routes for Gin Framework
	routes.Router(r)

	// Register Validators for Gin Framework
	validators.Register()

	port := os.Getenv("HTTP_PORT")
	r.Run(":" + port)

	fmt.Println("hais.in version:", os.Getenv("VERSION"))
	fmt.Println("Listening:", port)
}
