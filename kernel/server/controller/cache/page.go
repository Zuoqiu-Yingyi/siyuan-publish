package cache

import (
	"net/http"

	"publish/client"
	"publish/config"
	"publish/models"
	"publish/server/controller"
	"publish/server/status"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type page struct {
	controller.Page
}

var (
	P *page
)

func init() {
	P = &page{}
}

/* 解析 URL 参数 id */
func (*page) ID(c *gin.Context) {
	id := c.GetString("id")
	root_id := c.GetString("root_id")

	c.Request.URL.Path = "/block/" + root_id
	if id == root_id { // 查询文档块
		query := c.Request.URL.Query()
		query.Del("id")
		// fmt.Printf("%+v\n", query)
		// fmt.Printf("%+v\n", query.Encode())
		c.Request.URL.RawQuery = query.Encode()
		// router.HandleContext(c)
	}

	c.Redirect(http.StatusMovedPermanently, c.Request.URL.String())
}

/* 解析路由参数 id */
func (*page) Block(c *gin.Context) {
	root_id := c.GetString("id")
	doc := &models.Doc{}
	if d := doc.One(root_id); d == nil { // 数据库中没有该文档信息, 通过 API 查询
		var (
			r       *client.ResponseBody
			err     error
			err_msg string
		)

		/* 查询文档块 */
		r, err = client.GetBlockByID(client.C.R(), root_id)
		r, err_msg = client.Response(r, err)
		if r == nil {
			status.S.StatusInternalServerError(c, err_msg)
			return
		}
		doc_block := r.Data.([]interface{})
		switch {
		case len(doc_block) == 0:
			status.S.StatusBlockNotFound(c)
			return
		default:
			record := doc_block[0].(map[string]interface{})
			path := record["path"].(string)
			hpath := record["hpath"].(string)

			doc.ID = root_id
			doc.Path = string([]byte(path)[1 : len(path)-3])
			doc.Hpath = hpath
			doc.Title = record["content"].(string)
			doc.Tag = record["tag"].(string)

			// TODO 解析 IAL
			// block.Doc.Icon
			// block.Doc.TitleImg
		}

		/* 查询文档 DOM */
		r, err = client.GetBlockDomByID(client.C.R(), root_id, 0)
		r, err_msg = client.Response(r, err)
		if r == nil {
			status.S.StatusInternalServerError(c, err_msg)
			return
		}
		blocks := r.Data.(map[string]interface{})["blocks"].([]interface{})
		switch {
		case len(blocks) == 0:
			status.S.StatusBlockNotFound(c)
			return
		default:
			record := blocks[0].(map[string]interface{})
			doc.Dom = record["content"].(string)
		}
		models.DB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&doc)
	}

	c.HTML(http.StatusOK, "block.html", gin.H{
		"Path":     doc.Path,
		"Hpath":    doc.Hpath,
		"Tag":      doc.Tag,
		"Icon":     doc.Icon,
		"Title":    doc.Title,
		"TitleImg": doc.TitleImg,
		"Content":  doc.Dom,
		"Render":   config.C.Render,
	})
}
