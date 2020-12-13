package main

import (
	"light_blog/http_server"
	
)
func main() {
	fmt.Println("Service Start")
	//http_server.StartBlog()
	http_server.StartFileServer()
}
