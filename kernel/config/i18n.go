package config

/* 多语言相关设置 */
type I18n struct {
	Directory string `json:"directory"` // 多语言文件目录
	Default   string `json:"default"`   // 默认语言
}
