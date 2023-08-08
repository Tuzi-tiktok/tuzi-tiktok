// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID              int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:用户id" json:"id"`                      // 用户id
	Username        string         `gorm:"column:username;not null;comment:用户名" json:"username"`                                // 用户名
	Password        string         `gorm:"column:password;not null;comment:密码" json:"password"`                                 // 密码
	FollowCount     int64          `gorm:"column:follow_count;not null;comment:关注数" json:"follow_count"`                        // 关注数
	FollowerCount   int64          `gorm:"column:follower_count;not null;comment:粉丝数" json:"follower_count"`                    // 粉丝数
	Avatar          *string        `gorm:"column:avatar;comment:头像地址" json:"avatar"`                                            // 头像地址
	BackgroundImage *string        `gorm:"column:background_image;comment:用户个人页顶部大图" json:"background_image"`                   // 用户个人页顶部大图
	Signature       *string        `gorm:"column:signature;comment:个人签名" json:"signature"`                                      // 个人签名
	CreatedAt       time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt       *time.Time     `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`                                    // 更新时间
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间 支持软删除" json:"deleted_at"`                              // 删除时间 支持软删除
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
