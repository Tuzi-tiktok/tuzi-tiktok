package clients

import (
	"context"
	"fmt"
	"testing"
	"tuzi-tiktok/kitex/kitex_gen/favorite"
	"tuzi-tiktok/utils"
)

var ctx = context.TODO()

func TestGetFavorList(t *testing.T) {

	var req favorite.FavoriteListRequest
	req.UserId = int64(1)
	req.Token = "123"

	resp := new(favorite.FavoriteListResponse)

	cli, err := utils.NewFavorite()
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}

	resp, err = cli.GetFavoriteList(ctx, &req)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}

	fmt.Println(resp)
}

// 点赞操作
func TestFavorAction(t *testing.T) {
	var req favorite.FavoriteRequest
	req.Token = "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJQYXlsb2FkIjp7IlVJRCI6Nn0sImlzcyI6ImF1dGgtYXBpIiwiZXhwIjoxNjkzMDM3MzQ0fQ._44SDCTKSwBPfo0jXGH1ZE3INlDTxz2CqvEk_biOkDTfrxud9MTH-8iDy0mLYcqNm8P-Ksf5LbZb7ISXhNvWxQ"
	req.VideoId = 13
	req.ActionType = 1

	cli, err := utils.NewFavorite()
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	resp := new(favorite.FavoriteResponse)
	resp, err = cli.FavorVideo(ctx, &req)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(resp)
}

// 取消点赞操作
func TestUnFavorAction(t *testing.T) {
	var req favorite.FavoriteRequest
	req.Token = "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJQYXlsb2FkIjp7IlVJRCI6Nn0sImlzcyI6ImF1dGgtYXBpIiwiZXhwIjoxNjkzMDM3MzQ0fQ._44SDCTKSwBPfo0jXGH1ZE3INlDTxz2CqvEk_biOkDTfrxud9MTH-8iDy0mLYcqNm8P-Ksf5LbZb7ISXhNvWxQ"
	req.VideoId = 13
	req.ActionType = 2

	cli, err := utils.NewFavorite()
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	resp := new(favorite.FavoriteResponse)
	resp, err = cli.FavorVideo(ctx, &req)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(resp)
}
