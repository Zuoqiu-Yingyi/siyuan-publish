package client

import "publish/config"

/*
获得访问控制列表
	@return acl map[string]string: root_id => "public" | "protected" | "private"
	@return err_msg string: 错误消息
*/
func GetACL() (acl map[string]string, err_msg string) {
	var (
		r   *ResponseBody
		err error
	)
	/* 通过 API 查询具有访问控制字段的文档 ID 列表 */
	r, err = GetDocsByAttrName(C.R(), config.C.Siyuan.Publish.Access.Name)
	r, err_msg = Response(r, err)
	if r == nil {
		return
	}

	/* 建立文档 ID => 访问权限类型 public/protected/private 的映射 */
	data := r.Data.([]interface{})
	acl = make(map[string]string)
	for _, v := range data {
		record := v.(map[string]interface{})
		acl[record["root_id"].(string)] = record["value"].(string)
	}
	return
}
