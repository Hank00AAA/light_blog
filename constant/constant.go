package constant

import (
	"fmt"
	"light_blog/net_util"
)

const (
	HankShellURL = "http://127.0.0.1:8080/hankshell"
	FileParamKey = "file"
)

// GetURL
func GetURL() string {
	return fmt.Sprintf("http://%v:%v/hankshell", net_util.GetIP(), 8080)
}