package main

import (
	"flag"
	"time"

	"github.com/gin-gonic/gin"

	"github/leezone/simulator/common/db"
	"github/leezone/simulator/common/logger"
	"github/leezone/simulator/config"
	"github/leezone/simulator/middleware"
	"github/leezone/simulator/svc"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "path", "config/conf.toml", "config path")
	conf := config.NewConfig(configPath)

	logger.InitLogger(conf.Logger)
	db.InitSqlite(conf.Sqlite)
	svc.InitSvc()

	engine := gin.Default()
	engine.Use(middleware.Timeout)
	engine.GET("/ping", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)
		ctx.Writer.WriteString("success")
	})
	engine.Run(":9001")
}
