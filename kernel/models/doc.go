package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// 文档
type Doc struct {
	model
	ID       string `json:"id" gorm:"primaryKey;size:22"` // 文档块 ID
	Path     string `json:"path"`                         // 文档路径, 移除前缀 "/" 与 后缀 ".sy"
	Hpath    string `json:"hpath"`                        // 人类可读路径
	Tag      string `json:"tag"`                          // 文档标签
	Icon     string `json:"icon"`                         // 文档图标
	Title    string `json:"title"`                        // 文档标题
	TitleImg string `json:"title-img"`                    // 文档题头图
	Dom      string `json:"dom"`                          // 文档内容 DOM 字符串

	// 外键 => 访问控制信息
	Access Access `json:"access" gorm:"foreignKey:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (*Doc) Empty() bool {
	result := DB.Take(&Doc{})
	return errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func (doc *Doc) One(id string) interface{} {
	// REF [使用主键检索](https://gorm.io/zh_CN/docs/query.html#%E7%94%A8%E4%B8%BB%E9%94%AE%E6%A3%80%E7%B4%A2)
	doc.ID = id
	result := DB.First(&doc)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return doc
}
