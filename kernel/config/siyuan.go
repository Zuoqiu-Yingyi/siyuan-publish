package config

/* 思源相关设置 */
type Siyuan struct {
	Debug   bool    `json:"debug"`   // 是否开启调试模式
	Server  string  `json:"server"`  // 思源服务地址 http(s)://host:port
	Token   string  `json:"token"`   // 思源服务 token
	Timeout int     `json:"timeout"` // 向思源内核发起 HTTP 请求的超时时间(单位: ms)
	Retry   int     `json:"retry"`   // 向思源内核发起 HTTP 请求的重试次数
	Publish Publish `json:"publish"` // 发布相关设置
}

/* 思源发布相关设置 */
type Publish struct {
	Access Access `json:"access"` // 发布访问权限控制设置
}

/* 思源发布内容访问权限控制属性 */
type Access struct {
	Name      string    `json:"name"`      // 用于文档访问控制的属性名
	Default   string    `json:"default"`   // 默认访问权限
	Public    Attribute `json:"public"`    // 可公开访问的文档的属性(白名单)
	Protected Attribute `json:"protected"` // 鉴权后可公开访问的文档的属性
	Private   Attribute `json:"private"`   // 不可公开访问的文档的属性(黑名单)
}

/* 思源相关块属性 */
type Attribute struct {
	Name  string `json:"name"`  // 属性名称
	Value string `json:"value"` // 属性值
}
