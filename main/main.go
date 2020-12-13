package main

import (
	"fmt"
	"light_blog/http_server"
)

func main() {
	fmt.Println("Service Start")
	//http_server.StartBlog()
	http_server.StartFileServer()
}
