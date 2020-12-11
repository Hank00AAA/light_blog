package http_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// list_files
func list_files(ctx *gin.Context) {
	res := ""
	for _, v := range blogData.ListBlogFiles() {
		res = fmt.Sprintf("%v\n%v", res, v)
	}
	_, _ = ctx.Writer.Write([]byte(res))
}
