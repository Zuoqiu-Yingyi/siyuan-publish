package models

import (
	"errors"
	"math"
	"time"

	"publish/client"
	"publish/utils"

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

/*
查询文档信息并构造文档信息结构体
	@params id string: 文档 ID
	@return *Doc: 文档信息结构体
	@return error: 错误信息
*/
func GetDoc(id string) (*Doc, error) {
	var (
		r       *client.ResponseBody
		err     error
		err_msg string
		doc     = &Doc{}
	)

	/* 查询文档块 */
	r, err = client.GetBlockByID(client.C.R(), id)
	r, err_msg = client.Response(r, err)
	if r == nil {
		return nil, errors.New(err_msg)
	}
	doc_block := r.Data.([]interface{})
	switch {
	case len(doc_block) == 0:
		return nil, errors.New("document not found")

	default:
		record := doc_block[0].(map[string]interface{})
		path := record["path"].(string)
		hpath := record["hpath"].(string)

		doc.ID = id
		doc.Path = path[1 : len(path)-3]
		doc.Hpath = hpath[1:]
		doc.Title = record["content"].(string)
		doc.Tag = record["tag"].(string)

		ial := utils.IAL2Map(record["ial"].(string))
		doc.Icon = ial["icon"]
		doc.TitleImg = ial["title-img"]
	}

	/* 查询文档 DOM */

	// 使用嵌入块查询 API
	// r, err = client.GetBlockDomByID(client.C.R(), id, 0)
	// r, err_msg = client.Response(r, err)
	// if r == nil {
	// 	return nil, errors.New(err_msg)
	// }

	// blocks := r.Data.(map[string]interface{})["blocks"].([]interface{})
	// switch {
	// case len(blocks) == 0:
	// 	return nil, errors.New("document not found")
	// default:
	// 	record := blocks[0].(map[string]interface{})
	// 	doc.Dom = record["content"].(string)
	// }

	// 使用文档加载 API
	r, err = client.GetDoc(client.C.R(), id, "", 0, math.MaxInt32)
	r, err_msg = client.Response(r, err)
	if r == nil {
		return nil, errors.New(err_msg)
	}

	data := r.Data.(map[string]interface{})
	doc.Dom = data["content"].(string)
	return doc, nil
}
