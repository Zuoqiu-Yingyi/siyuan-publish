package config

/* 数据库相关设置 */
type Database struct {
	Debug  bool   `json:"debug"`  // 是否开启调试模式
	Reset  bool   `json:"reset"`  // 是否在每次启动时重置数据库
	SQLite string `json:"sqlite"` // 数据库文件路径
}
