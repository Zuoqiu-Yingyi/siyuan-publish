package client

import (
	"fmt"
	"net/http"

	"publish/config"

	"github.com/gin-gonic/gin"
	"github.com/imroc/req/v3"
)

/* 请求 */
func Request(request *req.Request, path string, body interface{}) (*ResponseBody, error) {
	r := &ResponseBody{}
	_, err := request.
		SetBodyJsonMarshal(body).
		SetResult(r).
		SetError(r).
		Post(config.C.Siyuan.Server + path)
	return r, err
}

/* SQL 查询 */
func SQL(request *req.Request, stmt string) (r *ResponseBody, err error) {
	body := struct {
		Stmt string `json:"stmt"`
	}{
		Stmt: stmt,
	}
	return Request(request, "/api/query/sql", body)
}

/* 通过 ID 或者指定块 */
func GetBlockByID(request *req.Request, id string) (r *ResponseBody, err error) {
	return SQL(request, fmt.Sprintf(`SELECT * FROM blocks WHERE id = '%s'`, id))
}

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

/* 通过 ID 或者指定块 */
func GetBlockDomByID(request *req.Request, id string) (r *ResponseBody, err error) {
	return SearchEmbedBlock(request, fmt.Sprintf(`SELECT * FROM blocks WHERE id = '%s'`, id), 0, make([]string, 0))
}

func Response(c *gin.Context, r *ResponseBody, err error) *ResponseBody {
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return nil
	}
	if r.Code != 0 {
		fmt.Println(r.Msg)
		c.String(http.StatusInternalServerError, r.Msg)
		return nil
	}
	return r
}
