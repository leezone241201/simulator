package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Logger struct {
	Type       string
	Level      string
	Encoding   string
	Path       string
	MaxSize    int  // 最大日志文件大小（MB），超过会切割
	MaxBackups int  // 保留的最大备份数
	MaxAge     int  // 保留的最大天数
	Compress   bool // 是否压缩旧日志
}

type Sqlite struct {
	Path string
}

type Config struct {
	Logger Logger
	Sqlite Sqlite
}

func NewConfig(path string) (config Config) {
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		log.Fatal("init config error ", err)
	}

	return
}
