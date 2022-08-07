package view

import (
	"net/http"
	"strings"

	"publish/config"
	"publish/models"

	"github.com/gin-gonic/gin"
)

func Block(c *gin.Context, doc *models.Doc) {
	var (
		title_height string   // 题头最小高度
		tags         []string // 文档标签
	)

	/* 文档题头 */
	if doc.TitleImg != "" { // 有题头图
		title_height = "30vh"
	} else if doc.Icon != "" { // 无题头图, 有图标
		title_height = "64px"
	} else { // 无题头图, 无图标
		title_height = "0"
	}

	/* 文档标签 */
	if doc.Tag != "" { // 有文档标签
		tags = strings.Split(doc.Tag[1:len(doc.Tag)-1], "# #")
	} else {
		tags = []string{}
	}

	c.HTML(http.StatusOK, "block.html", gin.H{
		"Home": gin.H{
			"URL":  config.C.Server.Index.URL,
			"Icon": config.C.Server.Index.Icon,
		},
		"Path":        doc.Path,
		"Hpath":       doc.Hpath,
		"Tags":        tags,
		"Icon":        doc.Icon,
		"Title":       doc.Title,
		"TitleImg":    doc.TitleImg,
		"TitleHeight": title_height,
		"Content":     doc.Dom,
		"Render":      config.C.Render,
	})
}
