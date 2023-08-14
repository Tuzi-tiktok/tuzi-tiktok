package dao

import (
	"context"
	"log"
	"testing"
	"tuzi-tiktok/dao/query"
)

var (
	qRelation = query.Relation
)

func TestIsFollower(t *testing.T) {
	followerID := 7
	followingID := 6

	count, err := qRelation.WithContext(context.Background()).Where(qRelation.FollowerID.Eq(int64(followerID)), qRelation.FollowingID.Eq(int64(followingID)), qRelation.DeletedAt.IsNull()).Count()
	if err != nil {
		return
	}

	log.Printf("count: %d", count)
}
