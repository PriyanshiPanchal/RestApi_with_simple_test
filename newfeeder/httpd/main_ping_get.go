package main

import (
	"github.com/gin-gonic/gin"
	"newsfeeder/httpd/demo"
	"newsfeeder/platform/newsfeed"
)

func main() {

	feed:=newsfeed.New()
	r:=gin.Default()

	r.GET("/ping", demo.PingGet)
	r.GET("/newsfeed", demo.NewsFeedGet(feed))
	r.POST("/newsfeed", demo.NewsFeedPost(feed))

	r.Run()
}
