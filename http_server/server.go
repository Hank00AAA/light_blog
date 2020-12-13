package http_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"light_blog/constant"
	"net/http"
	"os"
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
	err := http.ListenAndServe(":8080", HttpHandler{})
	if err != nil {
		panic(err)
	}
}

// httpHandler
type HttpHandler struct {

}

func (HttpHandler)ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	url := req.URL.String()
	fmt.Println(url)
	switch url {
	case "/webhook":
		attr := new(os.ProcAttr)
		newProcess, err := os.StartProcess("./webhook/webhook.sh", nil, attr)
		if err != nil {
			fmt.Println(err)
			_, _ = rsp.Write([]byte(err.Error()))
			return
		}
		fmt.Println("Process PID", newProcess.Pid)
		processState, err := newProcess.Wait() //等待命令执行完
		if err != nil {
			fmt.Println(err)
			_, _ = rsp.Write([]byte("webhook err!"))
			return
		}
		fmt.Println("processState PID:", processState.Pid())//获取PID
		fmt.Println("ProcessExit:", processState.Exited())//获取进程是否退出
		_, _ = rsp.Write([]byte("webhook finish!"))
	default:
		http.FileServer(http.Dir("./static_data")).ServeHTTP(rsp, req)
	}
}