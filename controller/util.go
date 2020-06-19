package controller

import (
	"github.com/aircjm/gocard/common"
	"github.com/aircjm/gocard/model/request"
	"github.com/aircjm/gocard/util"
	"github.com/gin-gonic/gin"
)

func Markdown2html(c *gin.Context) {
	request := request.Markdown2htmlRequest{}
	c.BindJSON(&request)
	cG := common.Gin{C: c}
	cG.Response(200, 0, util.ConvertMarkdown(request.MarkdownText))
}
