package http_server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"light_blog/constant"
	"net"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"time"
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
func StartFileServer(ln net.Listener) {
	var err error
	if ln == nil {
		l, err := net.Listen("tcp", ":80")
		if err != nil {
			panic(err)
		}
		ln = l
	} else {
		GetSingleTon().GracefulRestart()
	}

	GetSingleTon().SetListener(ln)
	svr := &http.Server{
		Handler:           HttpHandler{},
	}
	GetSingleTon().SetHttpServer(svr)
	fmt.Printf("pid:%v bindServer\n", GetSingleTon().GetPID())
	err = svr.Serve(ln)
	if err != nil {
		fmt.Println(fmt.Sprintf("pid:%v", GetSingleTon().GetPID()) + err.Error())
	}

	// 轮询链接，为0则关闭
	round := 0
	for {
		fmt.Println(fmt.Sprintf("pid:%v try close roud:%v curLink:%v links:%+v",
			GetSingleTon().GetPID(), round, GetSingleTon().GetLinkNum(),
			GetSingleTon().GetStatus()))
		round += 1
		time.Sleep(time.Second)
		curLink := GetSingleTon().GetLinkNum()
		if curLink == 0 {
			fmt.Println("Old Process exit pid:%v", GetSingleTon().GetPID())
			os.Exit(0)
		}
	}
}

// httpHandler
type HttpHandler struct {
}

func (HttpHandler) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	url := req.URL.String()
	GetSingleTon().AddLink(req.RemoteAddr)
	defer GetSingleTon().DelLink(req.RemoteAddr)
	fmt.Println(url)
	switch url {
	case "/webhook":
		if GetSingleTon().GetShutDown() {
			_, _ = rsp.Write([]byte("already process webhook"))
			return
		} else {
			GetSingleTon().ShutDown()
		}

		cmd := exec.Command("./webhook/webhook.sh")
		_, _ = rsp.Write([]byte("start"))
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Err:%v", err)
			return
		}
		fmt.Println(string(out))
		// restart
		//restart()
		// gracefulShutDown
		rsp.Write([]byte("has start shutdown"))
		go gracefulShutDown()
		return
	case "/keepalive":
		KeepAliveHandle(rsp, req)
	case "/shutdown":
		rsp.Write([]byte("shutdown"))
		os.Exit(0)
	default:
		http.FileServer(http.Dir("./static_data")).ServeHTTP(rsp, req)
	}
}

// gracefulShutDown
func gracefulShutDown() {
	// 将监听fd给
	forkChildProcess()
	fmt.Println("fork子进程监听FD ", GetSingleTon().GetPID())

	// 关闭服务端口
	time.Sleep(time.Second)
	err := GetSingleTon().GetHttpServer().Shutdown(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("关闭服务器 ", GetSingleTon().GetPID())
}

func forkChildProcess() {
	lisFD, err := GetSingleTon().GetListener().(*net.TCPListener).File()
	if err != nil {
		panic(err)
	}

	path := os.Args[0]
	fmt.Println("Path:%v", path)
	envList := []string{}
	for _, v := range os.Environ() {
		envList = append(envList, v)
	}

	execCall := &syscall.ProcAttr{
		Env:   envList,
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd(), lisFD.Fd()},
	}

	args := os.Args
	args = append(args, "-gr")

	fork, err := syscall.ForkExec(path, args, execCall)
	if err != nil {
		panic(err)
	}

	GetSingleTon().SetChildPid(fork)
	fmt.Printf("父进程:%v 子进程:%v\n", GetSingleTon().GetPID(), fork)
}

// restart
func restart() {
	binary, err := exec.LookPath("httpServer")
	if err != nil {
		panic(err)
	}

	args := []string{"httpServer"}

	env := os.Environ()

	if err := syscall.Exec(binary, args, env); err != nil {
		panic(err)
	}
}
