// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameVideo = "videos"

// Video mapped from table <videos>
type Video struct {
	ID            uint32         `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`                         // 主键
	AuthorID      uint32         `gorm:"column:author_id;not null;index:author_id,priority:1;comment:上传用户Id" json:"author_id"` // 上传用户Id
	Title         string         `gorm:"column:title;not null;comment:视频标题" json:"title"`                                      // 视频标题
	CoverURL      string         `gorm:"column:cover_url;not null;comment:封面url" json:"cover_url"`                             // 封面url
	PlayURL       string         `gorm:"column:play_url;not null;comment:视频播放url" json:"play_url"`                             // 视频播放url
	FavoriteCount uint32         `gorm:"column:favorite_count;not null;comment:点赞数" json:"favorite_count"`                     // 点赞数
	CommentCount  uint32         `gorm:"column:comment_count;not null;comment:评论数目" json:"comment_count"`                      // 评论数目
	CreatedAt     *time.Time     `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Video's table name
func (*Video) TableName() string {
	return TableNameVideo
}
