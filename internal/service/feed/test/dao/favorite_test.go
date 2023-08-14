package dao

import (
	"context"
	"log"
	"testing"
	"tuzi-tiktok/dao/query"
)

var qFavorite = query.Favorite

func TestIsFavorite(t *testing.T) {
	count, err := qFavorite.WithContext(context.Background()).Where(qFavorite.UID.Eq(1), qFavorite.Vid.Eq(3), qFavorite.DeletedAt.IsNull()).Count()
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
