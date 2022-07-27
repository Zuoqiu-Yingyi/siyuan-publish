package access

import (
	"strings"

	"publish/client"
	"publish/config"
	"publish/models"
	"publish/server/view/status"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func Cache(c *gin.Context) (bool, func(c *gin.Context), func(c *gin.Context)) {
	/* 判断 ACL 是否为空 */
	acl := &models.ACL{}
	if acl.Empty() {
		return false, status.S.StatusAccessDenied, nil
	}

	/* 判断数据库中是否有该块信息 */
	id := c.GetString("id")
	block := &models.Block{}
	if b := block.One(id); b == nil { // 数据库中没有该块信息, 通过 API 查询
		var (
			r       *client.ResponseBody
			err     error
			err_msg string
		)
		r, err = client.GetBlockByID(client.C.R(), id)
		r, err_msg = client.Response(r, err)
		if r == nil {
			return false, func(c *gin.Context) { status.S.StatusInternalServerError(c, err_msg) }, nil
		}
		data := r.Data.([]interface{})
		switch {
		case len(data) == 0:
			return false, status.S.StatusBlockNotFound, nil
		default:
			record := data[0].(map[string]interface{})
			root_id := record["root_id"].(string)
			path := record["path"].(string)

			block.ID = id
			block.RootId = root_id
			block.Path = path[1 : len(path)-3]

			models.DB.Create(block)
		}
	}

	c.Set("root_id", block.RootId)

	/* 判断数据库中是否有该块所在文档的访问权限信息 */
	permission := "?" // 初始化访问权限为未知权限
	access := &models.Access{}
	if a := access.One(block.RootId); a == nil { // 数据库中没有该块所在文档访问权限信息, 通过 API 查询并写入数据库
		docs := strings.Split(block.Path, "/")
		for i, j := 0, len(docs)-1; i < j; i, j = i+1, j-1 {
			docs[i], docs[j] = docs[j], docs[i]
		}
		records := make([]models.Access, 0, len(docs))
		for i, id := range docs {
			if a := access.One(id); a != nil { // Access 表中有 ID 为 id 的记录
				if ac := acl.One(access.ACL_ID); ac != nil { // ACL 表中有 ID 为 acl.ACL_ID 的记录
					permission = acl.Access
				} else { // 数据库逻辑错误
					return false, status.S.StatusPublishServerError, nil
				}
				for j := 0; j < i; j++ {
					records = append(records, models.Access{
						ID:     docs[j],
						ACL_ID: access.ACL_ID, // 从上级文档继承 ACL 项
					})
				}
				break
			}
		}
		if len(records) == 0 { // 没有便插入新的访问权限信息
			permission = config.C.Siyuan.Publish.Access.Default
			for _, id := range docs {
				records = append(records, models.Access{
					ID:     id,
					ACL_ID: "", // 关联 ID 为空的 ACL 项, 表示未知的访问权限
				})
			}
		}
		models.DB.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"acl_id", "updated_at"}),
		}).Create(&records) // 插入或更新
	} else { // 在数据库中查询到了文档的访问权限信息
		if a := acl.One(access.ACL_ID); a != nil { // ACL 表中有 ID 为 acl.ACL_ID 的记录
			permission = acl.Access
		} else { // 数据库逻辑错误
			return false, status.S.StatusPublishServerError, nil
		}
	}

	c.Set("access", permission)

	/* 根据访问权限判断是否继续或终止 */
	switch permission {
	case config.C.Siyuan.Publish.Access.Public.Value:
		return true, nil, nil
	case config.C.Siyuan.Publish.Access.Protected.Value:
		// TODO protected
		fallthrough
	case config.C.Siyuan.Publish.Access.Private.Value:
		fallthrough
	default:
		return false, status.S.StatusAccessDenied, nil
	}
}
