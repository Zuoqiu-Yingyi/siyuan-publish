package access

import (
	"publish/config"
	"publish/models"
	"publish/server/view/status"

	"github.com/gin-gonic/gin"
)

func Static(c *gin.Context) (bool, func(c *gin.Context), func(c *gin.Context)) {
	/* 判断 ACL 是否为空 */
	acl := &models.ACL{}
	if acl.Empty() {
		return false, status.S.StatusAccessDenied, nil
	}

	/* 判断数据库中是否有该块信息 */
	id := c.GetString("id")
	block := &models.Block{}
	if b := block.One(id); b == nil { // 数据库中没有该块
		return false, status.S.StatusBlockNotFound, nil
	}

	c.Set("root_id", block.RootId)

	/* 判断数据库中是否有该块所在文档的访问权限信息 */
	permission := "?" // 初始化访问权限为未知权限
	access := &models.Access{}
	if a := access.One(block.RootId); a == nil { // 数据库中没有该块所在文档访问权限信息
		return false, status.S.StatusBlockNotFound, nil
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
