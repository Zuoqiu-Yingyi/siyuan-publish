package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Block struct {
	model
	// [字段标签](https://gorm.io/zh_CN/docs/models.html#字段标签)
	ID   string `json:"id" gorm:"primaryKey;size:22"`
	Path string `json:"path"`

	// 外键
	// REF [Belongs To](https://gorm.io/zh_CN/docs/belongs_to.html)
	RootId string `json:"root_id" gorm:"size:22"`
	Access Access `json:"access" gorm:"foreignKey:RootId"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// REF [#TableName](https://gorm.io/zh_CN/docs/conventions.html#TableName)
// func (Block) TableName() string {
// 	return "blocks"
// }

func (*Block) Empty() bool {
	result := DB.Take(&Block{})
	return errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func (block *Block) One(id string) interface{} {
	// REF [使用主键检索](https://gorm.io/zh_CN/docs/query.html#%E7%94%A8%E4%B8%BB%E9%94%AE%E6%A3%80%E7%B4%A2)
	block.ID = id
	result := DB.First(&block)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return block
}
