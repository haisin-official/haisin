package database

import (
	"fmt"
	"log"

	"github.com/haisin-official/haisin/ent"
)

var (
	client *ent.Client
	err    error
)

func InitClient(driver string, dsn string) {
	client, err = ent.Open(driver, dsn)

	if err != nil {
		log.Panicf("Fatal Error of postgres 🚫\n %v", err)
	}
	fmt.Println("Success connection to postgres ✅")
}

func GetClient() *ent.Client {
	return client
}

func CloseClient() {
	client.Close()
}
