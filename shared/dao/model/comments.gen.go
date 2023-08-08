// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameComment = "comments"

// Comment mapped from table <comments>
type Comment struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:消息id" json:"id"`   // 消息id
	UID       int64          `gorm:"column:uid;not null;index:uid,priority:1;comment:用户id" json:"uid"` // 用户id
	Vid       int64          `gorm:"column:vid;not null;index:vid,priority:1;comment:视频id" json:"vid"` // 视频id
	Content   string         `gorm:"column:content;not null;comment:评论内容" json:"content"`              // 评论内容
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}
