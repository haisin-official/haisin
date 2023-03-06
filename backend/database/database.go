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

func InitDB(driver string, dsn string) {
	client, err = ent.Open(driver, dsn)

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	fmt.Println("Successful connect to postgres ✅")
}

func GetDB() *ent.Client {
	return client
}

func CloseDB() {
	client.Close()
}
