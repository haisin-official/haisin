package main

import (
	"fmt"

	"github.com/haisin-official/haisin/config"
)

func main() {
	fmt.Println("hais.in version:", config.GetEnv("VERSION"))
}
