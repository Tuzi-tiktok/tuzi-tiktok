package main

import (
	"context"
	"log"
	"time"
	f "tuzi-tiktok/kitex/kitex_gen/feed"
	feed "tuzi-tiktok/kitex/kitex_gen/feed/feedservice"
	"tuzi-tiktok/utils"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

func main() {
	c, err := feed.NewClient(utils.Feed(), client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		panic(err)
	}
	getOneVideo(c)
}

func getOneVideo(c feed.Client) {
	feedReq := &f.FeedRequest{
		LatestTime: nil,
		Token:      nil,
	}

	list, err := c.GetFeedList(context.Background(), feedReq, callopt.WithConnectTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(list)
}
