package main

import (
	"fmt"
	"newsfeeder/platform/newsfeed"
)

func main() {

	fmt.Println("Hello world")
	feed := newsfeed.New()

	fmt.Println(feed)

	feed.Add(newsfeed.Item{"Hello", "How ya' doing mate?"})

	fmt.Println(feed)
}
