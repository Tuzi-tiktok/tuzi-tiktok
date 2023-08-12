package test

import (
	"fmt"
	"testing"
	"tuzi-tiktok/dao/model"
	"tuzi-tiktok/dao/query"
)

func TestFavorVideo(t *testing.T) {
	f := query.Favorite
	uid := int64(1)
	vid := int64(10)
	favor := model.Favorite{UID: uid, Vid: vid}
	err := f.Create(&favor)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
}

func TestUnFavorVideo(t *testing.T) {
	f := query.Favorite
	uid := int64(1)
	vid := int64(10)
	result, err := f.Where(f.UID.Eq(uid), f.Vid.Eq(vid)).Delete()
	if result.RowsAffected == 0 {
		fmt.Println("Data does not exist. Deleting failed")
		t.Fail()
	}
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
}

func TestGetFavorList(t *testing.T) {
	f := query.Favorite
	uid := int64(1)
	videos, err := f.Where(f.UID.Eq(uid)).Find()
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	for _, v := range videos {

		fmt.Println(v.Vid)
	}
}
