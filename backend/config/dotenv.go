// 環境変数を読み込む
package config

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// production環境以外の場合, .envファイルから環境変数を読み込む
	if v := os.Getenv("GO_ENV"); v != "production" {
		err := godotenv.Load(".env")

		if err != nil {
			panic(".env is not found")
		}
	}
}

func GetEnv(name string) string {
	err := godotenv.Load(".env")

	if err != nil {
		panic(".env is not found!!")
	}

	v := os.Getenv(name)

	return v
}
