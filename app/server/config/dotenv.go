// 環境変数を読み込む
package config

import (
	"fmt"
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

		fmt.Println("Success load .env file ✅")
	}
}
