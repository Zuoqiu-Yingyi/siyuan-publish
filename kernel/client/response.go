package client

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
	@return string: 错误描述
*/
func Response(r *ResponseBody, err error) (*ResponseBody, string) {
	if err != nil {
		// fmt.Println(err)
		return nil, err.Error()
	}
	if r.Code != 0 {
		// fmt.Println(r.Msg)
		return nil, r.Msg
	}
	return r, ""
}
