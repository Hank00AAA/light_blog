package http_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"light_blog/constant"
	"os"
)

// hankShellHandleFunc
// 获取文件目录
func hankShellHandleFunc(ctx *gin.Context) {
	fileName := ctx.Param(constant.FileParamKey)
	content, err := readFile("./static_data/" + fileName)
	if err != nil {
		_, _ = ctx.Writer.Write(readFileWithoutErr("./static_data/index.html"))
		return
	}

	_, _ = ctx.Writer.Write(content)
}

// readFileWithoutErr
func readFileWithoutErr(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	return data
}

// readFile
func readFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil
}

// getAllFiles
func getAllFiles(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return files
}