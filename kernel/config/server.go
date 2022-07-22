package config

type Server struct {
	Debug     bool   `json:"debug"`     // 是否开启调试模式
	Port      int    `json:"port"`      // 服务端口
	Logs      string `json:"logs"`      // 日志目录
	Database  string `json:"database"`  // 数据库文件路径
	Templates string `json:"templates"` // 模板目录
	Mode      Mode   `json:"mode"`      // 模式
	Index     Index  `json:"index"`     // 首页配置
	Static    Static `json:"static"`    // 静态资源目录
}

type Mode struct {
	Page string `json:"page"` // 文档页面加载模式
	File string `json:"file"` // 资源文件加载模式
}

type Index struct {
	Paths []string `json:"path"` // 首页路径
	URL   string   `json:"url"`  // 首页重定向至的 URL
}

type Static struct {
	Appearance StaticPath `json:"appearance"` // 外观资源目录
	Assets     StaticPath `json:"assets"`     // 思源 assets 目录
	Emojis     StaticPath `json:"emojis"`     // 思源 emojis 目录
	Export     StaticPath `json:"export"`     // 思源 export 目录
	Stage      StaticPath `json:"stage"`      // 思源 stage 目录
	Widgets    StaticPath `json:"widgets"`    // 思源 widgets 目录
	JavaScript StaticPath `json:"javascript"` // javascript 目录
	CSS        StaticPath `json:"css"`        // css 目录
}
type StaticPath struct {
	Path     string `json:"path"`     // web 访问路径
	FilePath string `json:"filepath"` // 文件目录路径
}
