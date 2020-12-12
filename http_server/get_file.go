package http_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"light_blog/constant"
)

// handleGetFilesFunc
func handleGetFilesFunc(ctx *gin.Context) {
	fileName := ctx.Query(constant.FileParamKey)
	fmt.Println(fileName)
	content, err := blogData.GetFileContent(fileName)
	if err != nil {
		fmt.Printf("handleGetFilesFuncErr err:%v")
		_, _ = ctx.Writer.Write([]byte("err happen"))
		return
	}

	_, _ = ctx.Writer.Write([]byte(content))
}
