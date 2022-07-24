package models

import (
	"fmt"
	"publish/client"
	"publish/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	switch config.C.Server.Mode.Page {
	case "dynamic":
		fallthrough
	case "cache":
		if config.C.Database.Reset {
			clearDB() // 清空数据库
			initACL() // 初始化 ACL
		}
	case "static":
		if config.C.Database.Reset {
			clearDB()
			loadData()
		}
	default:
	}
}

/* 清空数据库所有记录 */
func clearDB() {
	/* 删除数据库中所有记录 */
	// REF [全局删除](https://gorm.io/zh_CN/docs/delete.html#%E9%98%BB%E6%AD%A2%E5%85%A8%E5%B1%80%E5%88%A0%E9%99%A4)
	global := DB.Session(&gorm.Session{AllowGlobalUpdate: true})
	global.Delete(&ACL{})
	global.Delete(&Access{})
	global.Delete(&Block{})
	global.Delete(&Doc{})
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

/* 载入数据 */
func loadData() {
	var (
		r       *client.ResponseBody
		err     error
		err_msg string
	)

	/* 查询所有 ACL 并按照 path 的长度升序排列 */
	r, err = client.SQL(client.C.R(), fmt.Sprintf(`
		SELECT
			a.root_id,
			a.value,
			a.path
		FROM
			attributes AS a
		WHERE
			a.name = '%s'
			AND a.root_id = a.block_id
		ORDER BY LENGTH(a.path);`,
		config.C.Siyuan.Publish.Access.Name,
	))
	r, err_msg = client.Response(r, err)
	if r == nil {
		panic(err_msg)
	}

	/* 构建 ACL 表 */
	data := r.Data.([]interface{})

	paths := make([]string, 0, len(data))
	acls := make([]ACL, 0, len(data))
	for _, v := range data {
		record := v.(map[string]interface{})
		path := record["path"].(string)

		paths = append(paths, path[:len(path)-3])
		acls = append(acls, ACL{
			ID:     record["root_id"].(string),
			Access: record["value"].(string),
		})
	}
	// TODO protected
	DB.Create(&acls) // 插入记录

	/* 按照顺序查询每条 ACL 记录对应的文档及其子文档并构建 Access 表 */
	for i := 0; i < len(paths); i++ {
		r, err = client.SQL(client.C.R(), fmt.Sprintf(`
			SELECT
				b.root_id
			FROM
				blocks AS b
			WHERE
				b.path LIKE '%s%%'
				AND b.type = 'd'
			ORDER BY LENGTH(b.path);`,
			paths[i],
		))
		r, err_msg = client.Response(r, err)
		if r == nil {
			panic(err_msg)
		}
		data = r.Data.([]interface{})

		accesses := make([]Access, 0, len(data))
		for _, v := range data {
			record := v.(map[string]interface{})

			accesses = append(accesses, Access{
				ID:     record["root_id"].(string),
				ACL_ID: acls[i].ID,
			})
		}
		DB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&accesses)
	}

	/* 为可访问的 Access 记录构建 Doc 表 */
	for _, acl := range acls {
		/* 获得可访问文档关联的 ACL 项 */
		switch acl.Access {
		case config.C.Siyuan.Publish.Access.Public.Value:
			fallthrough
		case config.C.Siyuan.Publish.Access.Protected.Value:
			accesses := make([]Access, 0, 0)
			/* 获得绑定某个可访问 ACL 项的所有文档 */
			DB.Where(&Access{ACL_ID: acl.ID}).Find(&accesses)
			for _, access := range accesses {
				/* 可访问文档的信息入库 */
				if doc, err := GetDoc(access.ID); err != nil {
					panic(err)
				} else {
					DB.Clauses(clause.OnConflict{
						UpdateAll: true,
					}).Create(doc)
				}
			}
		case config.C.Siyuan.Publish.Access.Private.Value:
			fallthrough
		default:
			continue
		}
	}

	/* 为可访问的 Access 记录构建 Block 表 */
	docs := make([]Doc, 0, 0)
	DB.Select("id", "path").Find(&docs)
	for _, doc := range docs {
		r, err = client.SQL(client.C.R(), fmt.Sprintf(`
			SELECT
				b.id
			FROM
				blocks AS b
			WHERE
				b.root_id = '%s';`,
			doc.ID,
		))
		r, err_msg = client.Response(r, err)
		if r == nil {
			panic(err_msg)
		}
		data = r.Data.([]interface{})

		blocks := make([]Block, 0, len(data))
		for _, v := range data {
			record := v.(map[string]interface{})

			blocks = append(blocks, Block{
				ID:     record["id"].(string),
				Path:   doc.Path,
				RootId: doc.ID,
			})
		}
		DB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&blocks)
	}
}
