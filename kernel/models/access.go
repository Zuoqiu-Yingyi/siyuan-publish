package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

/* 文档访问权限 */
type Access struct {
	model

	ID string `json:"id"`

	// 外键 => 访问控制列表
	ACL_ID string `json:"acl_id" gorm:"column:acl_id;size:22"`
	ACL    ACL    `json:"acl" gorm:"foreignKey:ACL_ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (*Access) Empty() bool {
	result := DB.Take(&Access{})
	return errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func (access *Access) One(id string) interface{} {
	// REF [使用主键检索](https://gorm.io/zh_CN/docs/query.html#%E7%94%A8%E4%B8%BB%E9%94%AE%E6%A3%80%E7%B4%A2)
	access.ID = id
	result := DB.First(&access)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return access
}
