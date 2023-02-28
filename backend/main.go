package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/haisin-official/haisin/config"
)

func main() {
	r := gin.Default()

	port := config.GetEnv("HTTP_PORT")
	r.Run(":" + port)

	fmt.Println("hais.in version:", config.GetEnv("VERSION"))
	fmt.Println("Listening:", port)

}
