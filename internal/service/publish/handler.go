package main

import (
	"context"
	"math/rand"
	"strings"
	"tuzi-tiktok/dao/model"
	"tuzi-tiktok/dao/query"
	publish "tuzi-tiktok/kitex/kitex_gen/publish"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/service/filetransfer/client"
	"tuzi-tiktok/utils/ffmpeg"
)

var transfer client.Transfer

func init() {
	transfer = client.NewTransfer()
}

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishVideo implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishVideo(ctx context.Context, req *publish.PublishRequest) (resp *publish.PublishResponse, err error) {
	resp = new(publish.PublishResponse)
	shot, err := ffmpeg.GetSnapShots(req.VideoUrl)
	if err != nil {
		return nil, err
	}
	pth := strings.Join([]string{"images", "1.png"}, "#")
	logger.Debug(pth)
	v := transfer.Put(pth, shot)
	if !v.Ok {
		resp.StatusCode = 400
		return
	}
	video := &model.Video{
		Title:    req.Title,
		AuthorID: rand.Uint32(),
		CoverURL: v.Url,
		PlayURL:  req.VideoUrl,
	}
	err = query.Video.Create(video)
	if err != nil {
		resp.StatusCode = 400
	}
	resp.StatusCode = 0
	return
}

// GetPublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) GetPublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	// TODO: Your code here...
	// Id
	//id := 1

	return
}
