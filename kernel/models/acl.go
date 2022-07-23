package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

/* 访问控制表(Access Control List) */
type ACL struct {
	model
	ID     string `json:"id" gorm:"primaryKey;size:22"` // 设置 custom-publish-access 属性的文档 ID
	Access string `json:"access"`                       // custom-publish-access 属性值(public/protected/private)

	CreatedAt time.Time
	UpdatedAt time.Time
}

/*
判断 ACL 表是否为空
REF [检索单个对象](https://gorm.io/zh_CN/docs/query.html#%E6%A3%80%E7%B4%A2%E5%8D%95%E4%B8%AA%E5%AF%B9%E8%B1%A1)
*/
func (*ACL) Empty() bool {
	// 获取一条记录，没有指定排序字段
	result := DB.Take(&ACL{})
	// SELECT * FROM users LIMIT 1;

	// 检查 ErrRecordNotFound 错误
	return errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func (acl *ACL) One(id string) interface{} {
	// REF [使用主键检索](https://gorm.io/zh_CN/docs/query.html#%E7%94%A8%E4%B8%BB%E9%94%AE%E6%A3%80%E7%B4%A2)
	acl.ID = id
	result := DB.First(&acl)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return acl
}
