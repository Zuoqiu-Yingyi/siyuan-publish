/*
 * 访问权限控制
 */

package auth

import (
	"strings"

	"publish/client"
	"publish/config"
	"publish/server/status"

	"github.com/gin-gonic/gin"
)

/*
判断文档是否可访问
REF [siyuan/session.go at master · siyuan-note/siyuan](https://github.com/siyuan-note/siyuan/blob/master/kernel/model/session.go)
*/
func Access(c *gin.Context) {
	var (
		r    *client.ResponseBody
		err  error
		data []interface{}
		path string
		docs []string
	)
	/* 通过 API 查询具有访问控制字段的文档 ID 列表 */
	r, err = client.GetDocsByAttrName(client.C.R(), config.C.Siyuan.Publish.Access.Name)
	r = client.Response(c, r, err)
	if r == nil {
		status.S.StatusSiyuanServerError(c)
		c.Abort()
		return
	}

	/* 建立文档 ID => 访问权限类型 public/protected/private 的映射 */
	data = r.Data.([]interface{})
	docs_with_access := make(map[string]string)
	switch {
	case len(data) == 0:
		status.S.StatusAccessDenied(c)
		c.Abort()
		return
	default:
		for _, v := range data {
			docs_with_access[v.(map[string]interface{})["root_id"].(string)] = v.(map[string]interface{})["value"].(string)
		}
	}

	/* 获得文档路径 */
	var id string // 查询的块 ID
	switch {
	case c.Query("id") != "":
		id = c.Query("id")
	case c.Param("id") != "":
		id = c.Param("id")
	default:
		status.S.StatusParamsError(c)
		c.Abort()
		return
	}
	r, err = client.GetBlockByID(client.C.R(), id)
	r = client.Response(c, r, err)
	if r == nil {
		status.S.StatusSiyuanServerError(c)
		c.Abort()
		return
	}
	data = r.Data.([]interface{})
	switch {
	case len(data) == 0:
		status.S.StatusBlockNotFound(c)
		c.Abort()
		return
	default:
		block := data[0].(map[string]interface{})
		path = block["path"].(string)
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
			c.Next()
			return
		case config.C.Siyuan.Publish.Access.Protected.Value:
			// TODO
			fallthrough
		case config.C.Siyuan.Publish.Access.Private.Value:
			status.S.StatusAccessDenied(c)
			c.Abort()
			return
		}
	}
	status.S.StatusAccessDenied(c)
	c.Abort()
}
