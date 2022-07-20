package client

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseBody struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

/*
解析响应体
	@params response *http.Response: HTTP 响应对象
	@params r *ResponseBody: 响应体
	@params err error: 错误

	@return *ResponseBody: 响应体
*/
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
