package http_server

import (
	"fmt"
	"light_blog/constant"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

// StartBlog
func StartBlog() {
	// peroidInit
	//peroidInit()

	// 注册函数
	r := gin.Default()
	//r.GET("/", hankShellHandleFunc)
	r.Any("/:"+constant.FileParamKey, hankShellHandleFunc)
	//r.GET("/hankshell/getFile", handleGetFilesFunc)

	// 允许http服务
	err := r.Run()
	if err != nil {
		panic(err)
	}
}

// StartFileServer
func StartFileServer() {
	err := http.ListenAndServe(":80", HttpHandler{})
	if err != nil {
		panic(err)
	}
}

// httpHandler
type HttpHandler struct {
}

func (HttpHandler) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	url := req.URL.String()
	fmt.Println(url)
	switch url {
	case "/webhook":
		cmd := exec.Command("./webhook/webhook.sh")
		_, _ = rsp.Write([]byte("start"))
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Err:%v", err)
			return
		}
		fmt.Println(string(out))
	default:
		http.FileServer(http.Dir("./static_data")).ServeHTTP(rsp, req)
	}
}
