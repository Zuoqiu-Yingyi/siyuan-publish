package client

import (
	"fmt"

	"github.com/imroc/req/v3"
)

/*
SQL 查询

	@params request *req.Request: HTTP 客户端请求对象
	@params stmt string: SQL 语句

	@return *ResponseBody: 响应体
	@return error: 错误
*/
func SQL(request *req.Request, stmt string) (r *ResponseBody, err error) {
	body := struct {
		Stmt string `json:"stmt"`
	}{
		Stmt: stmt,
	}
	return Request(request, "/api/query/sql", body)
}

/*
通过 ID 查询指定块的信息

	@params request *req.Request: HTTP 客户端请求对象
	@params id string: 块 ID

	@return *ResponseBody: 响应体
	@return error: 错误
*/
func GetBlockByID(request *req.Request, id string) (r *ResponseBody, err error) {
	return SQL(request, fmt.Sprintf(`SELECT * FROM blocks WHERE id = '%s' LIMIT 1;`, id))
}

/*
查询含有指定块属性名称的文档

	@params request *req.Request: HTTP 客户端请求对象
	@params attrName string: 属性名称

	@return *ResponseBody: 响应体
	@return error: 错误
*/
func GetDocsByAttrName(request *req.Request, attrName string) (r *ResponseBody, err error) {
	return SQL(request, fmt.Sprintf(`
		SELECT
			a.root_id,
			a.value
		FROM
			attributes AS a
		WHERE
			a.name = '%s'
			AND a.root_id = a.block_id;`,
		attrName,
	))
}

/*
通过嵌入块 API 获得块 DOM

	@params request *req.Request: HTTP 客户端请求对象
	@params stmt string: SQL 语句
	@params headingMode int: 对标题块的处理方式
		0: 同时显示标题块的下级块
		1: 仅显示标题块
	@params excludeIDs []string: 需要排除的块 ID

	@return *ResponseBody: 响应体
	@return error: 错误
*/
func SearchEmbedBlock(request *req.Request, stmt string, headingMode int, excludeIDs []string) (r *ResponseBody, err error) {
	body := struct {
		Stmt        string   `json:"stmt"`
		HeadingMode int      `json:"headingMode"`
		ExcludeIDs  []string `json:"excludeIDs"`
	}{
		Stmt:        stmt,
		HeadingMode: headingMode,
		ExcludeIDs:  excludeIDs,
	}
	return Request(request, "/api/search/searchEmbedBlock", body)
}

/*
查询文档

	@params request *req.Request: HTTP 客户端请求对象
	@params id string: 文档块/文档内顶级块 ID
	@params k string: 需要标记(高亮)的文本
	@params mode int: 查询模式
		REF https://github.com/siyuan-note/siyuan/blob/b94267be291ec0f2700c92ec210b4bde2462522f/kernel/model/file.go#L701
		0: 仅加载当前 ID 指定的块
		1: 加载上文
		2: 加载下文
		3: 加载上下文
		4: 从文档末尾加载上文
	@params size int: 查询的块数量

	@return *ResponseBody: 响应体
	@return error: 错误
*/
func GetDoc(request *req.Request, id, k string, mode, size int) (r *ResponseBody, err error) {
	body := struct {
		ID   string `json:"id"`
		K    string `json:"k"`
		Mode int    `json:"mode"`
		Size int    `json:"size"`
	}{
		ID:   id,
		K:    k,
		Mode: mode,
		Size: size,
	}
	return Request(request, "/api/filetree/getDoc", body)
}

/*
通过 ID 获得指定块的 DOM

	@params request *req.Request: HTTP 客户端请求对象
	@params id string: 块 ID
	@params headingMode int: 对标题块的处理方式
		0: 同时查询该标题块的下级块
		1: 仅查询该标题块

	@return *ResponseBody: 响应体
	@return error: 错误
*/
func GetBlockDomByID(request *req.Request, id string, headingMode int) (r *ResponseBody, err error) {
	return SearchEmbedBlock(request, fmt.Sprintf(`SELECT * FROM blocks WHERE id = '%s' LIMIT 1;`, id), headingMode, make([]string, 0))
}

/*
通过 pathname 使用 GET 方法获得指定资源

	@params request *req.Request: HTTP 客户端请求对象
	@params pathname string: URL pathname

	@return *Response: 响应体
	@return error: 错误
*/
func Get(request *req.Request, pathname string) (r *req.Response, err error) {
	r, err = request.
		Get(pathname)
	return
}
