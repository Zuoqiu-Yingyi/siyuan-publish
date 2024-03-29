package client

import (
	"github.com/imroc/req/v3"
)

/*
发起请求

	@params request *req.Request: HTTP 客户端请求对象
	@params pathname string: 请求 URL 的路径
	@params body interface{}: 请求体

	@return *ResponseBody: 响应体
	@return error: 错误
*/
func Request(request *req.Request, pathname string, body interface{}) (*ResponseBody, error) {
	r := &ResponseBody{}
	_, err := request.
		SetBodyJsonMarshal(body).
		SetResult(r).
		SetError(r).
		Post(pathname)
	return r, err
}
