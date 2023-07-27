package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const MysqlDsn = "root:admin@(localhost:3306)/tuzi_tiktok?charset=utf8mb4&parseTime=True&loc=Local"

type Querier interface {
}

func main() {

	// 连接数据库
	db, err := gorm.Open(mysql.Open(MysqlDsn))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:        "internal/dao/query",
		Mode:           gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
		FieldNullable:  true,  // generate pointer when field is nullable
		FieldCoverable: false, // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values

		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable:     true,  // detect integer field's unsigned type, adjust generated data type
		FieldWithIndexTag: true,  // generate with gorm index tag
		FieldWithTypeTag:  false, // generate with gorm column type tag
		WithUnitTest:      false,
	})
	// 设置目标 db
	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
	//  后续有需求 可生成动态SQL
	//g.ApplyInterface(func(querier Querier) {})

	g.Execute()

}
