package query

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	var err error
	db, err := gorm.Open(mysql.Open("root:admin@(localhost:3306)/tuzi_tiktok?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	SetDefault(db)
	fmt.Println("DataBase Init ")
}
