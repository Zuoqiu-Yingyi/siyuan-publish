package models

import "time"

// 文档
type Doc struct {
	ID       string `json:"id" gorm:"primaryKey;size:22"` // 文档块 ID
	Path     string `json:"path"`                         // 文档路径, 移除前缀 "/" 与 后缀 ".sy"
	Hpath    string `json:"hpath"`                        // 人类可读路径
	Icon     string `json:"icon"`                         // 文档图标
	Title    string `json:"title"`                        // 文档标题
	TitleImg string `json:"title-img"`                    // 文档题头图
	Dom      string `json:"dom"`                          // 文档内容 DOM 字符串

	// 外键 => 访问控制列表
	ACL_ID string `json:"acl_id" gorm:"column:acl_id;size:22"`
	ACL    ACL    `json:"acl" gorm:"foreignKey:ACL_ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
