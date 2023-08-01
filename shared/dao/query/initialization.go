package query

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	cfg "tuzi-tiktok/config"
	"tuzi-tiktok/logger"
)

func init() {

	dsn := cfg.DatabaseConfig.Dsn
	// TODO 数据库配置 GORM
	// TODO GORM日志整合 logger.Interface
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	SetDefault(db)
	logger.Info("- DataBase Init ")
}
