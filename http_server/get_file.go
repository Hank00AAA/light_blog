package http_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// handleGetFilesFunc
func handleGetFilesFunc(ctx *gin.Context) {
	fileName := ctx.Param("fileName")
	content, err := blogData.GetFileContent(fileName)
	if err != nil {
		fmt.Printf("handleGetFilesFuncErr err:%v")
		_, _ = ctx.Writer.Write([]byte("err happen"))
		return
	}

	_, _ = ctx.Writer.Write([]byte(content))
}
