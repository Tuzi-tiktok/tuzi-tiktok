package dao

import (
	"context"
	"log"
	"testing"
	"time"
	"tuzi-tiktok/dao/query"
)

var (
	qVideo = query.Video
)

func TestGetVideoWithTime(t *testing.T) {

	log.Println("Test getVideoWithTime")
	timeStr := "2023-05-05 07:06:15"

	tp, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		log.Printf("time.Parse err:  %v", err)
		return
	}

	log.Println(tp)

	videos, err := qVideo.WithContext(context.Background()).Where(qVideo.CreatedAt.Lt(tp), qVideo.DeletedAt.IsNull()).Order(qVideo.CreatedAt.Desc()).Limit(3).Find()

	if err != nil {
		return
	}

	log.Printf("length: %d", len(videos))
	nt := time.Now()
	for _, video := range videos {
		if nt.After(*video.CreatedAt) {
			nt = *video.CreatedAt
		}

		log.Printf("video = %v", video)
	}

	log.Printf("next_Time : %v", nt)

}

func TestCountVideos(t *testing.T) {
	count, err := qVideo.WithContext(context.Background()).Where(qVideo.AuthorID.Eq(1), qVideo.DeletedAt.IsNull()).Count()
	if err != nil {
		return
	}
	log.Printf("count: %d", count)
}

func TestGetTotalFavorite(t *testing.T) {
	// 获取作者作品vid列表
	vids, err := qVideo.WithContext(context.Background()).Select(qVideo.ID).Where(qVideo.AuthorID.Eq(1), qVideo.DeletedAt.IsNull()).Find()
	if err != nil {
		return
	}

	for _, vid := range vids {
		log.Printf("vid : %v", vid)
	}
}

func TestTime(t *testing.T) {
	var next int64
	next = 1655105759

	unix := time.Unix(next, 0)

	log.Println(unix)
}
