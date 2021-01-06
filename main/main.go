package main

import (
	"flag"
	"fmt"
	"light_blog/http_server"
	"net"
	"os"
)

var isGracefulRestart bool

func init() {
	flag.BoolVar(&isGracefulRestart, "gr", false, "优雅重启")
}

func main() {
	fmt.Println("Service Start")
	//http_server.StartBlog()
	flag.Parse()
	var l net.Listener
	if isGracefulRestart {
		fmt.Println("Graceful start")
		fd := os.NewFile(3, "")
		ln, err := net.FileListener(fd)
		if err != nil {
			panic(err)
		}
		l = ln
	} else {
		l = nil
	}
	http_server.StartFileServer(l)
}
