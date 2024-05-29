package models

import (
	"bluebell/pkg/sonyflake"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	Id        uint64         `json:"id" gorm:"primary_key; column:id; type:bigint(20) unsigned; not null;"`
	ParentId  uint64         `json:"parentId" gorm:"column:parent_id; type:bigint(20)"`
	PostId    uint64         `json:"postId" gorm:"column:post_id; type:bigint(20)"`
	AuthorId  uint64         `json:"authorId" gorm:"column:author_id; type:bigint(20)"`
	Content   string         `json:"content" gorm:"column:content; type:longtext; not null;"`
	CreatedAt time.Time      `json:"-" gorm:"column:created_at; autoCreateTime;"`
	UpdatedAt time.Time      `json:"-" gorm:"column:updated_at; autoUpdateTime;"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type CommentInfo struct {
	Id       uint64 `json:"id" form:"id" binding:"omitempty"`
	ParentId uint64 `json:"parentId" form:"parentId" binding:"omitempty"`
	PostId   uint64 `json:"postId" form:"postId" binding:"omitempty"`
	AuthorId uint64 `json:"authorId" form:"authorId" binding:"omitempty"`
	Content  string `json:"content" form:"content" binding:"omitempty"`
}

type CommentIds struct {
	Ids []uint64 `json:"ids"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := sonyflake.GetId()
	if err != nil {
		return
	}
	c.Id = id
	return
}
