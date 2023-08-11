package query

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	cfg "tuzi-tiktok/config"
	"tuzi-tiktok/logger"
)

var DatabaseConfig databaseConfig

type databaseConfig struct {
	Dsn          string // 数据来源名称
	Username     string // 用户
	Password     string // 密码
	Host         string // 数据库链接
	Port         uint64 // 数据库端口
	DataBaseName string // 数据库名
	Timeout      string // 连接超时
}

const databaseK = "database"

func config() {
	log.Println("- Load DataSource Config")
	if err := cfg.VConfig.GetViper().UnmarshalKey(databaseK, &DatabaseConfig); err != nil {
		panic(err)
	}
	// Check Dsn Validity
	if DatabaseConfig.Dsn == "" {
		log.Println(" - Init Dsn")
		DatabaseConfig.Dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
			DatabaseConfig.Username,
			DatabaseConfig.Password,
			DatabaseConfig.Host,
			DatabaseConfig.Port,
			DatabaseConfig.DataBaseName,
			DatabaseConfig.Timeout,
		)
	}
}

func init() {
	// TODO GORM日志整合 logger.Interface

	config()
	db, err := gorm.Open(mysql.Open(DatabaseConfig.Dsn), &gorm.Config{})
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	SetDefault(db)
}
