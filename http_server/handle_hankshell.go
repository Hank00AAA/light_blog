package http_server

import (
	"github.com/gin-gonic/gin"
)

// hankShellHandleFunc
func hankShellHandleFunc(ctx *gin.Context) {
	list_files(ctx)
}
