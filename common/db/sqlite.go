package db

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github/leezone/simulator/common/logger"
	"github/leezone/simulator/config"
)

var SqliteDB *gorm.DB

func InitSqlite(conf config.Sqlite) {
	var err error
	SqliteDB, err = gorm.Open(sqlite.Open(conf.Path))
	if err != nil {
		logger.Logger.ErrorWithStack("init sqlite error", err, conf)
		os.Exit(-1)
	}
	logger.Logger.Info("init sqlite success")
}
