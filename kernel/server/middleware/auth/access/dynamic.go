package access

import (
	"strings"

	"publish/client"
	"publish/config"
	"publish/server/status"

	"github.com/gin-gonic/gin"
)

func Dynamic(c *gin.Context) (bool, func(c *gin.Context), func(c *gin.Context)) {
	var (
		r    *client.ResponseBody
		err  error
		data []interface{}
		path string
		docs []string
	)

	/* 获得访问控制列表 root_id => "public" | "protected" | "private" */
	acl, err_msg := client.GetACL()
	if acl == nil {
		return false, func(c *gin.Context) { status.S.StatusInternalServerError(c, err_msg) }, nil
	}
	if len(acl) == 0 {
		return false, status.S.StatusAccessDenied, nil
	}

	/* 获得文档路径 */
	id := c.GetString("id")
	r, err = client.GetBlockByID(client.C.R(), id)
	r, err_msg = client.Response(r, err)
	if r == nil {
		return false, func(c *gin.Context) { status.S.StatusInternalServerError(c, err_msg) }, nil
	}
	data = r.Data.([]interface{})
	switch {
	case len(data) == 0:
		return false, status.S.StatusBlockNotFound, nil
	default:
		block := data[0].(map[string]interface{})
		path = block["path"].(string)
		c.Set("root_id", block["root_id"].(string))
	}

	/* 将文档路径分割为文档 ID 列表, 按照层级从深到浅排列 */
	docs = strings.Split(path[1:len(path)-3], "/")
	for i, j := 0, len(docs)-1; i < j; i, j = i+1, j-1 {
		docs[i], docs[j] = docs[j], docs[i]
	}
	// fmt.Printf("%#v\n%#v\n", acl, docs)

	/* 根据访问权限判断是否继续或终止 */
	for _, root_id := range docs {
		if right, ok := acl[root_id]; ok {
			c.Set("access", right)
			switch right {
			case config.C.Siyuan.Publish.Access.Public.Value:
				return true, nil, nil
			case config.C.Siyuan.Publish.Access.Protected.Value:
				// TODO
				fallthrough
			case config.C.Siyuan.Publish.Access.Private.Value:
				fallthrough
			default:
				return false, status.S.StatusAccessDenied, nil
			}
		}
	}
	c.Set("access", "?")
	return false, status.S.StatusAccessDenied, nil
}
