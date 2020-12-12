package http_server

import (
	"github.com/gin-gonic/gin"
	"light_blog/constant"
	"net/http"
)

// StartBlog
func StartBlog() {
	// peroidInit
	//peroidInit()

	// 注册函数
	r := gin.Default()
	//r.GET("/", hankShellHandleFunc)
	r.Any("/:" + constant.FileParamKey, hankShellHandleFunc)
	//r.GET("/hankshell/getFile", handleGetFilesFunc)

	// 允许http服务
	err := r.Run()
	if err != nil {
		panic(err)
	}
}

// StartFileServer
func StartFileServer() {
	err := http.ListenAndServe(":8080", http.FileServer(http.Dir("./")))
	if err != nil {
		panic(err)
	}
}