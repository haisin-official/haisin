package session

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/haisin-official/haisin/app/utils"
	goRedis "github.com/redis/go-redis/v9"
)

type Manager struct {
	MaxAge   int
	Path     string
	Domain   string
	Secure   bool
	HttpOnly bool
}

// Cookieが持つ値
type Store struct {
	SessionId string `json:"session_id" binding:"required"`
	UserId    string `json:"user_id" binding:"required"`
}

var (
	r       *goRedis.Client
	manager Manager
)

func init() {
	url := os.Getenv("REDIS_URL")
	opt, err := goRedis.ParseURL(url)
	if err != nil {
		log.Panicf("Fatal Error of redis 🚫\n %v", err)
	}
	r = goRedis.NewClient(opt)

	fmt.Println("Success connection to redis ✅")

	manager.MaxAge = 86400
	manager.Path = "/"
	manager.Domain = os.Getenv("HTTP_HOST")
	manager.Secure, _ = strconv.ParseBool(os.Getenv("SESSION_SECURE"))
	manager.HttpOnly = true
}

func NewSession(c *gin.Context, cKey string, v Store) {
	// Generate New Token
	token := utils.GenRandToken()

	vj, _ := json.Marshal(v)

	if err := r.Set(c, token, vj, 0).Err(); err != nil {
		log.Panicf("Fatal Error of Register Session 🚫 \n %v", err)
	}

	fmt.Println("new session 😎", cKey, token)
	c.SetCookie(cKey, token, manager.MaxAge, manager.Path, manager.Domain, manager.Secure, manager.HttpOnly)
}

// セッションの値を取得
func GetSession(c *gin.Context, cKey string) (*Store, int, error) {
	token, _ := c.Cookie(cKey)
	vJSON, err := r.Get(c, token).Bytes()
	if err == goRedis.Nil {
		fmt.Println("No Registered Token ⚠")
		return nil, http.StatusUnauthorized, fmt.Errorf("no registered token %v", err)
	}
	if err != nil {
		fmt.Println("Error has occured in getting Session 🚫")
		return nil, http.StatusInternalServerError, fmt.Errorf("error has occured in getting session %v", err)
	}

	// Convert value to store struct
	v := new(Store)
	if err := json.Unmarshal(vJSON, v); err != nil {
		fmt.Println("Error has occured in Unmarshal JSON 🚫")
		return nil, http.StatusInternalServerError, fmt.Errorf("error has occured in unmarshal json %v", err)
	}
	return v, http.StatusOK, nil
}

// セッションを削除
func DeleteSession(c *gin.Context, cKey string) error {
	token, _ := c.Cookie(cKey)
	r.Del(c, token)
	c.SetCookie(cKey, "", -1, manager.Path, manager.Domain, manager.Secure, manager.HttpOnly)
	return nil
}
