package server

import (
	"net/http"

	"github.com/go-redis/redis/v7"
	"golang.org/x/net/context"
)

type AuthService struct {
	cacheClient redis.ClusterClient
}

func NewAuthServer(ctx *context.Context, redisEndPointAddr []string, password string) *AuthService {
	return &AuthService{cacheClient: redis.NewClusterClient(&redis.ClusterOptions{Addrs: redisEndPointAddr, PoolSize: 5, Password: password})}
}

func setupHTTPRoutePath() {
	http.HandleFunc("auth", func(rw http.ResponseWriter, r *http.Request) {

	})
	http.HandleFunc("signin", func(rw http.ResponseWriter, r *http.Request) {

	})
	http.ListenAndServe(":3000", nil)
}
