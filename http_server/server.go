package http_server

import (
	"github.com/gin-gonic/gin"
)

// StartBlog
func StartBlog() {
	// peroidInit
	peroidInit()

	// 注册函数
	r := gin.Default()
	r.GET("/hankshell", hankShellHandleFunc)
	r.GET("/hankshell/getFile/:fileName", handleGetFilesFunc)

	// 允许http服务
	err := r.Run()
	if err != nil {
		panic(err)
	}
}