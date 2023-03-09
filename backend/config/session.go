package config

import (
	"net/http"
	"strconv"
	"time"

	"github.com/alexedwards/scs/goredisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var (
	sessionManager *scs.SessionManager
)

func NewSessions(client *redis.Client) {
	// Init manager and use goredisstore
	sessionManager = scs.New()
	sessionManager.Store = goredisstore.New(client)

	// Manager Sessings
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Name = "haisin_session"
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteStrictMode

	sessionManager.Cookie.Domain = GetEnv("HTTP_HOST")
	secureSecret, _ := strconv.ParseBool(GetEnv("SESSION_SECURE"))
	sessionManager.Cookie.Secure = secureSecret
}

func GetSessionManager() *scs.SessionManager {
	return sessionManager
}

func GetSessionValue(ctx *gin.Context, key string) string {
	msg := sessionManager.GetString(ctx, key)
	return msg
}

func SetSessionValue(ctx *gin.Context, key string, value string) {
	sessionManager.Put(ctx, key, value)
}
