package models

import (
	"publish/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() {
	/* 链接数据库 */
	// REF [连接到数据库 | GORM](https://gorm.io/zh_CN/docs/connecting_to_the_database.html#SQLite)
	// file::memory:?cache=shared
	if db, err := gorm.Open(sqlite.Open(config.C.Server.Database), &gorm.Config{}); err != nil {
		panic(err)
	} else {
		DB = db
	}

	/* 迁移数据库 */
	// REF [迁移 | GORM](https://gorm.io/zh_CN/docs/migration.html)
	DB.AutoMigrate(&ACL{}, &Doc{}, &Block{})

	/* 初始化 ACL */
	initACL()
}

func initACL() {

}
