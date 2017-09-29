package main

import (
	"flag"

	"github.com/sundy-li/gos/gofetcher"
)

var (
	urlStr string
)

func init() {
	flag.StringVar(&urlStr, "url", "http://www.baidu.com", "input the url")
}
func main() {
	flag.Parse()
	spider := gofetcher.NewFetcher(urlStr)
	spider.AddIgnore(".mp4")
	err := spider.Execute()
	if err != nil {
		println(err.Error())
	}
}
