package http_server

import (
	"fmt"
	"github.com/robfig/cron"
	"light_blog/blog_data"
	"light_blog/git_files"
)

var blogData blog_data.BlogData

// peroidFunc
func peroidInit() {
	updateBlogData()

	crontab := cron.New()
	err := crontab.AddFunc("*/60  * * * * ", updateBlogData)
	if err != nil {
		fmt.Printf("cron cronTabErr:%v", err)
		return
	}
	crontab.Start()
}

// updateBlogData
func updateBlogData() {
	r := git_files.UpdateLocalCache()
	if r == nil {
		return
	}

	blogData.Repo = r
}