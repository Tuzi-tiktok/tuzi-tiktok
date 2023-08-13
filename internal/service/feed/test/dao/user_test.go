package dao

import (
	"context"
	"log"
	"testing"
	"tuzi-tiktok/dao/query"
)

var (
	qUser = query.User
)

func TestGetUserInfo(t *testing.T) {

	var aid int64
	aid = 1

	user, err := qUser.WithContext(context.Background()).Where(qUser.ID.Eq(aid), qUser).First()
	if err != nil {
		panic(err)
	}

	log.Printf("user : %v", user)

}
