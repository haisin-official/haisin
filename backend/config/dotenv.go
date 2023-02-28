// 環境変数を読み込む
package config

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(name string) string {
	err := godotenv.Load(".env")

	if err != nil {
		panic(".env is not found!!")
	}

	v := os.Getenv(name)

	return v
}
