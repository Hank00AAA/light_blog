package http_server

import (
	"github.com/gin-gonic/gin"
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
	m.MarkdownAddTitle(1, "小可爱下午英语考试加油～～～～")
	m.MarkdownAddEmptyRow()

	//for _, v := range blogData.ListBlogFiles() {
	//	m.MarkdownHttp(v, fmt.Sprintf("%v/getFile?" + constant.FileParamKey + "=%v", constant.GetURL(), v))
	//	m.MarkdownAddEmptyRow()
	//}

	m.AddPic("miao", "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1607752834804&di=b416c06b3cf91c4e402041e0b8f302f1&imgtype=0&src=http%3A%2F%2Fb-ssl.duitang.com%2Fuploads%2Fitem%2F201803%2F08%2F20180308223406_zKyVN.jpeg")
}