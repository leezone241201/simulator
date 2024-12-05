package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.Writer.WriteString("success")
	})
}
