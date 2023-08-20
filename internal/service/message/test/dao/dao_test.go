package dao_test

import (
	"context"
	"log"
	"testing"
	"time"
	"tuzi-tiktok/dao/model"
	"tuzi-tiktok/dao/query"
	"tuzi-tiktok/redis"
)

var qMessage = query.Message

func TestDao(t *testing.T) {
	tKey := "15-16-pre-message-time"
	v, err := redis.IRC.Get(context.Background(), tKey).Result()
	if err != nil {
		panic(err)
	}
	log.Println(v)

	toUid := 15
	fromUid := 16
	loc, err := time.LoadLocation("Asia/Shanghai")
	pt, _ := time.ParseInLocation("2006-01-02 15:04:05", v, loc)
	log.Printf("pre - time is %v", pt)

	mList, _ := qMessage.WithContext(context.Background()).Where(qMessage.ToUserID.Eq(int64(toUid))).
		Where(qMessage.FormUserID.Eq(int64(fromUid))).
		Where(qMessage.CreatedAt.Gt(pt)).
		Find()
	if len(mList) == 0 {
		log.Printf(" m list is nil")
	}

	printMList(mList)
}

func printMList(ml []*model.Message) {
	for _, m := range ml {
		log.Printf("m--> %v", m)
	}
}
