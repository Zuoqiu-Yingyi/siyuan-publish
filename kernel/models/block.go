package models

import "time"

type Block struct {
	// [字段标签](https://gorm.io/zh_CN/docs/models.html#字段标签)
	ID string `json:"id" gorm:"primaryKey;size:22"`

	// 外键
	// REF [Belongs To](https://gorm.io/zh_CN/docs/belongs_to.html)
	RootId string `json:"root_id" gorm:"size:22"`
	Doc    Doc    `gorm:"foreignKey:RootId"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// REF [#TableName](https://gorm.io/zh_CN/docs/conventions.html#TableName)
// func (Block) TableName() string {
// 	return "blocks"
// }
