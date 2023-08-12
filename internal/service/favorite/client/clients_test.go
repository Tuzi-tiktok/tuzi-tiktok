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

func TestFavorAction(t *testing.T) {
	var req favorite.FavoriteRequest
	req.Token = "123"
	req.VideoId = 11
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

func TestUnFavorAction(t *testing.T) {
	var req favorite.FavoriteRequest
	req.Token = "123"
	req.VideoId = 11
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
