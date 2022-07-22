package models

import "time"

/* 访问控制表(Access Control List) */
type ACL struct {
	ID     string `json:"id" gorm:"primaryKey;size:22"` // 设置 custom-publish-access 属性的文档 ID
	Access string `json:"access"`                       // custom-publish-access 属性值(public/protected/private)

	CreatedAt time.Time
	UpdatedAt time.Time
}
