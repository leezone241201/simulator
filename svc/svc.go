package svc

import (
	"github/leezone/simulator/common/db"
	"github/leezone/simulator/common/logger"
	"github/leezone/simulator/dao"
)

var Svc ServerContext

type ServerContext struct {
	FileDB           dao.FileDB
	ServerTemplateDB dao.ServerTemplateDB
}

func InitSvc() {
	Svc = ServerContext{
		FileDB:           dao.NewFileDB(db.SqliteDB),
		ServerTemplateDB: dao.NewServerTemplateDB(db.SqliteDB),
	}

	Svc.FileDB.AutoMigrate()
	Svc.FileDB.AutoMigrate()

	logger.Logger.Info("init server context success")
}
