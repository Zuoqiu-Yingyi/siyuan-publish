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
		r       *client.ResponseBody
		err     error
		err_msg string
		data    []interface{}
		path    string
		docs    []string
	)
	/* 通过 API 查询具有访问控制字段的文档 ID 列表 */
	r, err = client.GetDocsByAttrName(client.C.R(), config.C.Siyuan.Publish.Access.Name)
	r, err_msg = client.Response(r, err)
	if r == nil {
		return false, func(c *gin.Context) { status.S.StatusInternalServerError(c, err_msg) }, nil
	}

	/* 建立文档 ID => 访问权限类型 public/protected/private 的映射 */
	data = r.Data.([]interface{})
	docs_with_access := make(map[string]string)
	switch {
	case len(data) == 0:
		return false, status.S.StatusAccessDenied, nil
	default:
		for _, v := range data {
			docs_with_access[v.(map[string]interface{})["root_id"].(string)] = v.(map[string]interface{})["value"].(string)
		}
	}

	/* 获得文档路径 */
	id := c.GetString("id")
	if id == "" {
		return false, status.S.StatusParamsError, nil
	}
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
	docs = strings.Split(string([]byte(path)[1:len(path)-3]), "/")
	for i, j := 0, len(docs)-1; i < j; i, j = i+1, j-1 {
		docs[i], docs[j] = docs[j], docs[i]
	}
	// fmt.Printf("%#v\n%#v\n", docs_with_access, docs)
	for _, doc_id := range docs {
		switch docs_with_access[doc_id] {
		case config.C.Siyuan.Publish.Access.Public.Value:
			c.Set("access", config.C.Siyuan.Publish.Access.Public.Value)
			return true, nil, nil
		case config.C.Siyuan.Publish.Access.Protected.Value:
			c.Set("access", config.C.Siyuan.Publish.Access.Protected.Value)
			// TODO
			fallthrough
		case config.C.Siyuan.Publish.Access.Private.Value:
			c.Set("access", config.C.Siyuan.Publish.Access.Private.Value)
			return false, status.S.StatusAccessDenied, nil
		}
	}
	return false, status.S.StatusAccessDenied, nil
}
