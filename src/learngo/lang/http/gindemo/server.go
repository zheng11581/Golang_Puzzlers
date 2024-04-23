package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"time"
)

var requestIdKey string = "requestId"

func main() {
	r := gin.New()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	r.Use(func(context *gin.Context) {
		start := time.Now()
		context.Next()
		// log path, status, latency
		logger.Info("incoming request",
			zap.String("path", context.Request.URL.Path),
			zap.Int("status", context.Writer.Status()),
			zap.Duration("duration", time.Now().Sub(start)),
		)
		context.Next()
	}, func(context *gin.Context) {
		context.Set(requestIdKey, rand.Int())
		context.Next()
	},
	)

	r.GET("/ping", func(context *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if rid, exists := context.Get(requestIdKey); exists {
			h[requestIdKey] = rid
		}

		context.JSON(http.StatusOK, h)
	})
	r.GET("/hello", func(context *gin.Context) {
		h := gin.H{
			"message": "hello",
		}
		if rid, exists := context.Get(requestIdKey); exists {
			h[requestIdKey] = rid
		}
		context.JSON(http.StatusOK, h)
	})
	r.Run(":8080")
}
