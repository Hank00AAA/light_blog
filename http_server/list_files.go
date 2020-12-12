package http_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"light_blog/constant"
	"light_blog/markdown"

	//"light_blog/markdown"
)

// list_files
func list_files(ctx *gin.Context) {

	ctx.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	m := markdown.NewMarkdown()
	createMarkDownFileList(m)
	_, _ = ctx.Writer.Write(m.MarkdownToHtml())
}

// createMarkDownList
func createMarkDownFileList(m *markdown.MarkDownData) {
	m.MarkdownAddTitle(1, "HanlShell")
	m.MarkdownAddEmptyRow()

	for _, v := range blogData.ListBlogFiles() {
		m.MarkdownHttp(v, fmt.Sprintf("%v/getFile?" + constant.FileParamKey + "=%v", constant.GetURL, v))
		m.MarkdownAddEmptyRow()
	}

	m.AddPic("miao", "https://github.com/Hank00AAA/light_blog/blob/main/pic/miao.jpeg?raw=true")
}