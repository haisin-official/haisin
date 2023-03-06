package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/haisin-official/haisin/config"
	"github.com/haisin-official/haisin/routes"
)

func main() {
	r := gin.Default()
	routes.Router(r)

	port := config.GetEnv("HTTP_PORT")
	r.Run(":" + port)

	fmt.Println("hais.in version:", config.GetEnv("VERSION"))
	fmt.Println("Listening:", port)
}
