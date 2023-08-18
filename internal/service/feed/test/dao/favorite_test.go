package dao

import (
	"context"
	"log"
	"testing"
	"tuzi-tiktok/dao/query"
)

var qFavorite = query.Favorite

func TestIsFavorite(t *testing.T) {
	var uid int64
	var vid int64
	uid = 4
	vid = 1
	dInfo, _ := qFavorite.WithContext(context.Background()).
		Where(qFavorite.UID.Eq(uid)).
		// Where(qFavorite.Vid.Eq(vid)).
		Delete()
	log.Printf("delete %v ", dInfo)

	count, err := qFavorite.WithContext(context.Background()).Where(qFavorite.UID.Eq(uid), qFavorite.Vid.Eq(vid), qFavorite.DeletedAt.IsNull()).Count()
	if err != nil {
		return
	}

	log.Printf("count %d", count)
}

func TestGetUserFavorite(t *testing.T) {
	count, err := qFavorite.WithContext(context.Background()).Where(qFavorite.UID.Eq(6), qFavorite.DeletedAt.IsNull()).Count()
	if err != nil {
		return
	}

	log.Printf("count %d", count)
}
