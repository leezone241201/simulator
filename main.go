package main

import (
	"github/leezone/simulator/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.Use(middleware.Timeout)
	engine.GET("/ping", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)
		ctx.Writer.WriteString("success")
	})
	engine.Run(":9001")
}
