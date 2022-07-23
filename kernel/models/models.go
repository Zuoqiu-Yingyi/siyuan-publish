package models

import (
	"publish/client"
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
	if db, err := gorm.Open(sqlite.Open(config.C.Database.SQLite), &gorm.Config{}); err != nil {
		panic(err)
	} else {
		DB = db
	}

	/* 迁移数据库 */
	// REF [迁移 | GORM](https://gorm.io/zh_CN/docs/migration.html)
	DB.AutoMigrate(&ACL{}, &Access{}, &Block{}, &Doc{})

	if config.C.Database.Reset {
		/* 删除数据库中所有记录 */
		// REF [全局删除](https://gorm.io/zh_CN/docs/delete.html#%E9%98%BB%E6%AD%A2%E5%85%A8%E5%B1%80%E5%88%A0%E9%99%A4)
		global := DB.Session(&gorm.Session{AllowGlobalUpdate: true})
		global.Delete(&ACL{})
		global.Delete(&Access{})
		global.Delete(&Block{})
		global.Delete(&Doc{})

		/* 初始化 ACL */
		initACL()
	}
}

/* 初始化 ACL*/
func initACL() {
	/* 获得访问控制列表 root_id => "public" | "protected" | "private" */
	acl, err_msg := client.GetACL()
	if acl == nil {
		panic(err_msg)
	}

	/* 插入或更新记录 */
	acls := make([]ACL, 0, len(acl)+1)
	accesses := make([]Access, 0, len(acl))

	acls = append(acls, ACL{ID: "", Access: "?"}) // 空记录表示未设置访问权限
	for root_id, access := range acl {
		acls = append(acls, ACL{
			ID:     root_id,
			Access: access,
		})
		accesses = append(accesses, Access{
			ID:     root_id,
			ACL_ID: root_id,
		})
	}

	// REF [批量插入](https://gorm.io/zh_CN/docs/create.html#批量插入)
	DB.Create(&acls)     // 插入记录
	DB.Create(&accesses) // 插入记录

	// REF [Upsert 及冲突](https://gorm.io/zh_CN/docs/create.html#Upsert-%E5%8F%8A%E5%86%B2%E7%AA%81)
	// DB.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "id"}},
	// 	DoUpdates: clause.AssignmentColumns([]string{"access", "updated_at"}),
	// }).Create(&ACLs) // 插入或更新
}
